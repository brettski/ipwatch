# ipwatch

![version](https://img.shields.io/badge/version-3.1.0-blue)  
![nodejs](https://img.shields.io/badge/nodejs->=12-darkgreen)  
![hello](https://img.shields.io/badge/hi-👋-lightgray)  

ipWatch keeps an eye on my external IP. If it changes, then send an email alert and optionally update CloudFlare DNS.

It's all pretty basic, really.

There are two pieces for ipWatch: The ipWatch app and a server (a simple cloud function) to send requests to get the requester's IP address.

## The GCP function

This simple function is set up in GCP and pointed to by the server.js script. It returns the server headers as a JSON object.

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

## Using Cloudflare Workers

I discovered recently that Cloudflare will provide a request IP header value on the free plan. Good news for our little app here. Here's a worker script that will work with ipWatch.

```js
var worker_default = {
  async fetch(request, env, ctx) {
    const ipHeaderValue = 'cf-connecting-ip';
    const output = {
      stamp: new Date(),
      'x-real-ip': request.headers.get('x-real-ip'),
      'x-forwarded-for': request.headers.get(ipHeaderValue),
    }
    return Response.json(output);
  }
};
export {
  worker_default as default
};
```

## JavaScript Function

On each run, the function examines data from the endpoint and searches the database for the `x-forwarded-for` header value.

- if found, add one to seen count
- if not found, send email and add new record.
- _new in v3_, update CloudFlare DNS with the new IP record
  - If the updates status is indicated in ip change email notification

It uses nedb and shouldn't grow large as it will only have as many records as IPs discovered. Which hopefully is very few, or there are bigger issues with your ISP changing IPs all the time.

## THANK YOU

Thank you to these projects for their great work! I appreciate you.

- [node-fetch](https://www.npmjs.com/package/node-fetch)
- [debug](https://www.npmjs.com/package/debug)
- [dotenv](https://www.npmjs.com/package/dotenv)
- [nedb](https://www.npmjs.com/package/nedb) && [nedb-promises](https://www.npmjs.com/package/nedb-promises)
- [postmark](https://postmarkapp.com), [node client lib](https://www.npmjs.com/package/postmark)

**Service bulletin from me:** Get off SendGrid, use [Postmark](https://postmarkapp.com), they are so much better. Yeah, it's more of a hassle to get started with, but to me, that's worth it as it keeps the bad players off their service. Go check them out.

## Installation

In a nutshell, it's all manual. A little embarrassing seeing that I have spent much of my career automating all the things.

### Environment

```sh
ENDPOINT_CHK=https://example.com
POSTMARK_TOKEN=--token-value--
DEBUG=brettski:*
EMAIL_TO=
EMAIL_FROM=
```

Variables used for Cloudflare update

```sh
UPDATE_ON_CHANGE=false 
# CloudFlare Zone Updates
CF_BEARER_TOKEN=
CF_ZONE_DOMAIN=
CF_DNS_RECORD_HOSTNAME=
```

`ENDPOINT_CHK` is the endpoint of the function code. There is no file for the function at this time as it's really straight forward. Simply copy `GCP Function` listed above into a gcp function, Azure function etc. CloudFlare workers can't be used for this as they don't expose client IP's (the call it real client IPs) unless you have an enterprise account. And well I'm not an enterprise with enterprise pockets.

`server.js` can be run anywhere there is node 12+ and write access to it's folder for the db file. This needs to run from somewhere within your network, calling out to the internet so the function can relay back your IP for recording.

I have this running on my home Synology using a scheduled task to execute the node file. It does the job.

## CloudFlare DNS Updates

By getting a CloudFlare API token for your account update the domain and verify when the isp domain changes. This functionality if very beta (not well tested)

```sh
UPDATE_ON_CHANGE=true 
# CloudFlare Zone Updates
CF_BEARER_TOKEN=
CF_ZONE_DOMAIN=
CF_DNS_RECORD_DOMAIN=
```

`UPDATE_ON_CHANGE` a feature switch to turn updates on or off.

`CF_BEARER_TOKEN` the API token from a CloudFlare account at [https://dash.cloudflare.com/profile/api-tokens](https://dash.cloudflare.com/profile/api-tokens).

`CF_ZONE_DOMAIN` the domain of the zone we're updating, how the domain is in the list of domains managed in CloudFlare. This is **NOT** the full domain name. Something like, `brettski.com` not `house.brettski.com`.

`CF_DNS_RECORD_HOSTNAME` is the full domain we are updating with the new IP address.

Yes, naming is hard.

## In Use

So earlier this our ISP sent a 'reset' code to fix something. It changed the IP being used, and an email was sent as hoped. 😅 We were able to change DNS references to keep cross-location services running.
