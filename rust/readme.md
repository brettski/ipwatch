# ipWatch Rust

## ⚠️ WIP ⚠️

The Rust version of ipWatch is a wip and no where close to a working application.

Whats needed:

- Finish error trapping and reporting in requests
- Structure files
- Set up requests to send emails
- Find a local, file-based db to use for tracking IP's. (e.g. anydb in JS)

So yeah, a lot.

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

Yes, I'd like to write the cloud function in Rust too, but most services don't support Rust yet. Once the service is done I'll look closer at some offerings and put them here.

## THANK YOU

Thank you to the following crates allowing this application to be possible.

- [dotenv](https://github.com/dotenv-rs/dotenv), [@crates.io](https://crates.io/crates/dotenv)
- [reqwest](https://github.com/seanmonstar/reqwest), [@crates.io](https://crates.io/crates/reqwest)
- [Postmark](https://github.com/pastjean/postmark-rs), [@crates.io](https://crates.io/crates/postmark)
- [clap](https://github.com/clap-rs/clap), [@crates.io](https://crates.io/crates/clap)
