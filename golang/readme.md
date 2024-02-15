# ipWatch (golang)

The Golang version of ipWatch.

![version](https://img.shields.io/badge/version-1.0.0-lightblue)
![golang](https://img.shields.io/badge/golang-%3E=1.20-lightblue)
![hello](https://img.shields.io/badge/hi-👋-lightgray)

ipWatch keeps an eye on my external IP. If it changes then sends an email alert and optionally update CloudFlare DNS.

All pretty basic really.

There are two pieces for ipWatch. The ipWatch app and a server (simple cloud function) to send requests to for getting the current IP address.

## Command Line Commands

use -h for help menus

## Remote Function

See [nodejs readme](../nodejs/readme.md) for setting up remote function.

## THANK YOU

Thank you to these projects for their great work! I appreciate you.

<!-- - [requests](https://requests.readthedocs.io/en/latest/) -->
- [mow.cli](https://github.com/jawher/mow.cli)  
- [godotenv](https://github.com/joho/godotenv)  
- [bbolt](https://go.etcd.io/bbolt)  

## Installation

Set the following environment variables by using a shell script or a .env file. There is a sample (.env.sample) as a reference.

```sh
ENDPOINT_CHK=https://example.com
POSTMARK_TOKEN=--token-value--
EMAIL_TO=
EMAIL_FROM=
```

**Notice** CloudFlare update not supported in Golang version, yet.

There will be executables for macOS, Linux (Flavors TBD), and 64-bit Windows
