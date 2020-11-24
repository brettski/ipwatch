# ipwatch

It keeps an eye on my external IP. If it changes then an alert is sent.

All pretty basic really.

## The GCP function

```JavaScript
/**
 * Responds to any HTTP request.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */
exports.headerDump = (req, res) => {
  const head = JSON.stringify(req.headers)
  delete head['if-none-match']
  delete head['x-appengine-default-version-hostname']
  res.status(200).send(head)
};
```

## function

On every run looks at data from endpoint and searches for the x-forwarded-for header value in the database.

- if found, add one to seen
- if not found, send email add new record.

Uses nedb and shouldn't grow large as it will only have as many records as IP's discovered. Which hopefully is very few or there are bigger issues with your crap ISP. 

## THANK YOU

Thank you to these projects for their great work! I appreciate you.

- [axios](https://www.npmjs.com/package/axios)
- [debug](https://www.npmjs.com/package/debug)
- [dotenv](https://www.npmjs.com/package/dotenv)
- [moment](https://www.npmjs.com/package/moment)
- [nedb](https://www.npmjs.com/package/nedb) & [nedb-promise](https://www.npmjs.com/package/nedb-promise)
- [postmark](https://postmarkapp.com), [node client lib](https://www.npmjs.com/package/postmark)

Get off SendGrid, use [Postmark](https://postmarkapp.com), they are so much better. Yeah, more of a hassle to get started with, but to me that's worth it as it keeps the bad players off their service. Go check them out if you have a moment.

## Installation

In a nutshell it's all manual. A little embarrassing seeing that I have spent much of my career automating all the things.

### Environment

```sh
ENDPOINT_CHK=https://example.com
POSTMARK_TOKEN=--token-value--
DEBUG=brettski:*
EMAIL_TO=
EMAIL_FROM=
```

Using Postmark for email. I really like their service and they care about deliverability of their clients. Check them out if you have a moment.

`ENDPOINT_CHK` is the endpoint of the function code. There is no file for the function at this time as it's really straight forward. Simply apply `GCP Function` listed above into a gcp function or even a CloudFlare worker.

`server.js` can be run anywhere there is node 10+ and write access to it's folder for the db file. This needs to run from somewhere within your network, calling out to the internet so the function can relay back your IP for recording. 

I have this running on my home Synology using a scheduled task to execute the node file. It does the job. 

## In Use

So earlier this our ISP sent a 'reset' code to fix something. It changed the IP being used and an email was sent as hoped. 😅 We were able to change references to keep cross location services running. 
