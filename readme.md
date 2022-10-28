# ipwatch

It keeps an eye on my external IP. If it changes then an alert is sent.

All pretty basic really.

## The GCP function

This simple function is set up in GCP and pointed to by the server.js script. It returns the server headers as a json object.

```JavaScript
/**
 * Responds to any HTTP request.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */
exports.headerDump = (req, res) => {
  const head = req.headers
  delete head['if-none-match']
  delete head['x-appengine-default-version-hostname']
  res.status(200).json(head)
};
```

## function

On each run the function looks at data from endpoint and searches for the `x-forwarded-for` header value in the database.

- if found, add one to seen count
- if not found, send email and add new record.

Uses nedb and shouldn't grow large as it will only have as many records as IP's discovered. Which hopefully is very few or there are bigger issues with your crap ISP.

## THANK YOU

Thank you to these projects for their great work! I appreciate you.

- [node-fetch](https://www.npmjs.com/package/node-fetch)
- [debug](https://www.npmjs.com/package/debug)
- [dotenv](https://www.npmjs.com/package/dotenv)
- [nedb](https://www.npmjs.com/package/nedb) & [nedb-promises](https://www.npmjs.com/package/nedb-promises)
- [postmark](https://postmarkapp.com), [node client lib](https://www.npmjs.com/package/postmark)

Get off SendGrid, use [Postmark](https://postmarkapp.com), they are so much better. Yeah, more of a hassle to get started with, but to me that's worth it as it keeps the bad players off their service. Go check them out.

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

`ENDPOINT_CHK` is the endpoint of the function code. There is no file for the function at this time as it's really straight forward. Simply copy `GCP Function` listed above into a gcp function, Azure function etc. CloudFlare workers can't be used for this as they don't expose client IP's (the call it real client IPs) unless you have an enterprise account. And well I'm not an enterprise with enterprise pockets. 

`server.js` can be run anywhere there is node 12+ and write access to it's folder for the db file. This needs to run from somewhere within your network, calling out to the internet so the function can relay back your IP for recording.

I have this running on my home Synology using a scheduled task to execute the node file. It does the job.

## In Use

So earlier this our ISP sent a 'reset' code to fix something. It changed the IP being used and an email was sent as hoped. 😅 We were able to change DNS references to keep cross location services running.

Perhaps a next version can update the DNS record in Cloudflare with the changed IP value. 🤔
