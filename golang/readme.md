# ipWatch (python)

The python version of ipWatch.

![version](https://img.shields.io/badge/version-0.1.0-lightblue)
![golang](https://img.shields.io/badge/%3E=1.22-lightblue)
![hello](https://img.shields.io/badge/hi-👋-lightgray)

ipWatch keeps an eye on my external IP. If it changes then sends an email alert and optionally update CloudFlare DNS.

All pretty basic really.

There are two pieces for ipWatch. The ipWatch app and a server (simple cloud function) to send requests to for getting the current IP address.

## Command Line Commands

use -h for help menus

## Notice

Hopefully future me has provided compiled executables for the latest version

## THANK YOU

Thank you to these projects for their great work! I appreciate you.

- TBD
<!-- - [requests](https://requests.readthedocs.io/en/latest/) -->

## Installation

In a nutshell it's all manual. A little embarrassing seeing that I have spent much of my career automating all the things.

```sh
ENDPOINT_CHK=https://example.com
POSTMARK_TOKEN=--token-value--
EMAIL_TO=
EMAIL_FROM=
```

**Notice** CloudFlare update not supported in Golang version yet.

There will be executables for macOS, Linux (Flavors TBD), and 64-bit Windows
