require('dotenv').config();
const Datastore = require('nedb-promise')
const axios = require('axios');
const moment = require('moment');
const postmark = require('postmark');
const debug = require('debug');

dlog = debug('brettski:ipwatch');

const endpoint = process.env.ENDPOINT_CHK;
const postmarktoken = process.env.POSTMARK_TOKEN;
const emailto = process.env.EMAIL_TO;
const emailfrom = process.env.EMAIL_FROM;
const client = new postmark.Client(postmarktoken);

const db = Datastore({
  filename: './db.data',
  autoload: true,
  timestampData: true,
})

axios.get(endpoint)
  .then(response => {
    let curip = '';
    if (response && response.data && response.data['x-forwarded-for']) {
      curip = response.data['x-forwarded-for'];
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
            ip: response.data['x-forwarded-for'],
            seen: 1,
          }).then(d => dlog('insterted %o', d));
        }
      })
    }
    client.sendEmail({
      From: emailfrom,
      To: emailto,
      Subject: `No repsponse: from endpoint ipwatch`,
      TextBody: `No response for ip check at ${moment().format()}\nendpoint: ${endpoint}`,
    })
    dlog('this %o', response.data['x-forwarded-for'])
  })
  .catch((error) => {
    console.error('that shit wasn\'t right', error);
  })
  