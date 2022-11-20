![version](https://img.shields.io/badge/version-0.1.0-red) 
![hello](https://img.shields.io/badge/hi-👋-lightgray)


ipWatch keeps an eye on my external IP. If it changes then sends an email alert ~~and optionally update CloudFlare DNS~~.

All pretty basic really.

There are two pieces for ipWatch. The ipWatch app and a server (simple cloud function) to send requests to for getting the current IP address.

## Rust-lang app

Rust version of ipWatch which provides and executable to run for the check. The author knows very little rust so be forewarned!

## The GCP function

This simple function is set up in GCP and pointed to by the server.js script. It returns the server headers as a json object.

This is the cloud function used by the [nodejs](../nodejs/readme.md) version of ipWatch.

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
