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
exports.headDump = (req, res) => {
  const head = JSON.stringify(req.headers)
  res.status(200).send(head);
};
```

## function

On every run looks at data from endpoint and searches for the x-forwarded-for header value in the database.

- if found, add one to seen
- if not found, send email add new record.

Uses nedb and shouldn't grow large as it will only have as many records as IP's discovered. Which hopefully is very few or there are bigger issues with your crap ISP. 

## THANK YOU

Thank you to these projects for thier great work! I appreciate you.

- [axios](https://www.npmjs.com/package/axios)
- [debug](https://www.npmjs.com/package/debug)
- [dotenv](https://www.npmjs.com/package/dotenv)
- [moment](https://www.npmjs.com/package/moment)
- [nedb](https://www.npmjs.com/package/nedb) & [nedb-promise](https://www.npmjs.com/package/nedb-promise)
- [postmark](https://postmarkapp.com), [node client lib](https://www.npmjs.com/package/postmark)
