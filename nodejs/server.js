require('dotenv').config();
const Datastore = require('nedb-promises');
const fetch = require('node-fetch');
const postmark = require('postmark');
const debug = require('debug');
const envConfig = require('./envConfig');
const { cloudFlare, validateDns } = require('./lib/cloudFlare');

const dlog = debug('brettski:ipwatch');

const { endpoint, testLocation, updateOnChange } = envConfig;
const postmarktoken = envConfig.postmark.token;
const emailto = envConfig.postmark.emailto;
const { zoneDomain, recordHostname } = envConfig.CloudFlare;
const emailfrom = envConfig.postmark.emailfrom;
const client = new postmark.Client(postmarktoken);

const db = Datastore.create({
  filename: './db.data',
  autoload: true,
  timestampData: true,
});

fetch(endpoint)
  .then(res => res.json())
  .then(response => {
    let curip = '';
    if (response && response['x-forwarded-for']) {
      curip = response['x-forwarded-for'];
      db.findOne({
        ip: curip,
      }).then(async result => {
        if (result) {
          dlog('found ip %s; seen: %d times.', curip, result.seen);
          db.update(
            {
              _id: result._id,
            },
            {
              $set: {
                seen: (result.seen += 1),
              },
            },
          ).then(up => console.log('db update', up));
        } else {
          // new IP address
          // send email and add to database
          dlog('New IP discovered! %s', curip);
          let body = `A new IP address has been discoved at ${testLocation}, ${curip}`;
          if (
            zoneDomain !== null &&
            recordHostname !== null &&
            updateOnChange === true
          ) {
            // update ip in CloudFlare, get result add to message body
            const cf = await cloudFlare();
            const {id: zoneId} = await cf.getZoneId(zoneDomain);
            const dnsRecord = await cf.getDnsRecordId(zoneId, recordHostname);
            dlog('dnsRecord: %o', dnsRecord);
            const { id: dnsRecordId } = dnsRecord;
            const updateStatus = await cf.updateDnsRecordIp({
              zoneId,
              dnsRecordId,
              newIp: curip,
            });
            // updated, probably check status here
            dlog('update ip status', updateStatus);
            // validate new dns record.
            const isNewIpSet = validateDns(recordHostname, curip);
            if (isNewIpSet) {
              body += '\n\n IP successfully updated at CloudFlare';
            } else {
              body +=
                '\n\n IP DNS validation check came back false. Manually validate new IP is set in DNS';
            }
          } else {
            body +=
              '\n\n**DNS NOT updated**. Updates are disabled or not configured. The new IP will need to be updated manually.';
          }
          // yeah this is a Promise, but we don't care about the result so waiting
          // for a response from the service doesn't matter.
          client.sendEmail({
            From: emailfrom,
            To: emailto,
            Subject: 'New IP discovered from our isp',
            TextBody: body,
          });
          db.insert({
            timestamp: new Date().toISOString(),
            ip: response['x-forwarded-for'],
            seen: 1,
          }).then(d => {
            dlog('insterted %o', d);
            console.log('inserted db', d);
          });
        }
      });
    } else {
      client.sendEmail({
        From: emailfrom,
        To: emailto,
        Subject: `No repsponse: from endpoint ipwatch`,
        TextBody: `No response for ip check at ${new Date().toString()}\nendpoint: ${endpoint}`,
      });
    }
    dlog('this %o', response['x-forwarded-for']);
  })
  .catch(error => {
    console.error("that shit wasn't right", error);
  });
