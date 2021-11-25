require('dotenv').config();
const Datastore = require('nedb-promises')
const fetch = require('node-fetch');
const moment = require('moment');
const postmark = require('postmark');
const debug = require('debug');
const envConfig = require('./envConfig');

dlog = debug('brettski:ipwatch');

const endpoint = envConfig.endpoint;
const postmarktoken = envConfig.postmark.token;
const emailto = envConfig.postmark.emailto;
const emailfrom = envConfig.postmark.emailfrom;
const client = new postmark.Client(postmarktoken);

const db = Datastore.create({
  filename: './db.data',
  autoload: true,
  timestampData: true,
})

fetch(endpoint)
  .then(res => res.json())
  .then(response => {
    let curip = '';
    if (response && response['x-forwarded-for']) {
      curip = response['x-forwarded-for'];
      db.findOne({
        ip: curip,
      }). then (result => {
        if (result) {
          dlog('found ip %s; seen: %d times.', curip, result.seen);
          db.update({
            _id: result._id}, 
            { $set: {
              seen: result.seen += 1
              }
            }
          ).then(up => console.log('upper', up));
        } else {
          // new IP address
          // send email and add to database
          dlog('New IP discovered! %s', curip);
          const body = `A new IP address has been discoved at Brookmeade. You know what to do \n\nip: ${curip}`;
          client.sendEmail({
            From: emailfrom,
            To: emailto,
            Subject: "New IP from our isp",
            TextBody: body,
          })
          db.insert({
            timestamp: moment().format(),
            ip: response['x-forwarded-for'],
            seen: 1,
          }).then(d => dlog('insterted %o', d));
        }
      })
    } else {
      client.sendEmail({
        From: emailfrom,
        To: emailto,
        Subject: `No repsponse: from endpoint ipwatch`,
        TextBody: `No response for ip check at ${moment().format()}\nendpoint: ${endpoint}`,
      })
    }
    dlog('this %o', response['x-forwarded-for'])
  })
  .catch((error) => {
    console.error('that shit wasn\'t right', error);
  })
  