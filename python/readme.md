# ipWatch (python)

The python version of ipWatch.

![version](https://img.shields.io/badge/version-1.0.0-blue)
![python](https://img.shields.io/badge/nodejs->=3.9-3776AB)
![hello](https://img.shields.io/badge/hi-👋-lightgray)

ipWatch keeps an eye on my external IP. If it changes then sends an email alert and optionally update CloudFlare DNS.

All pretty basic really.

There are two pieces for ipWatch. The ipWatch app and a server (simple cloud function) to send requests to for getting the current IP address.

## Command Line Commands

Let's build some command line options. Here are some thoughts

- run the test (ip check)
- list current ip from db
- list all ip from db

## Notice

See [nodejs readme](../nodejs/readme.md) for setting up remote function.

## THANK YOU

Thank you to these projects for their great work! I appreciate you.

- [python-dotenv](https://pypi.org/project/python-dotenv/)
- [tinydb](https://tinydb.readthedocs.io/en/latest/)
- [postmarker](https://postmarkapp.com/send-email/python), [postmark](https://postmarkapp.com)
- [requests](https://requests.readthedocs.io/en/latest/)

## Installation

In a nutshell it's all manual. A little embarrassing seeing that I have spent much of my career automating all the things.

```sh
ENDPOINT_CHK=https://example.com
POSTMARK_TOKEN=--token-value--
DEBUG=brettski:*
EMAIL_TO=
EMAIL_FROM=
```

**Notice** CloudFlare update not supported in Python version yet.

`ipwatch.py` is the main file.

**TO DO**: write python/pip installation instructions.  

First setup a [venv](https://docs.python.org/3/library/venv.html) to run things in

then, you at least need to do `pip install -r ./requirements.txt`.
