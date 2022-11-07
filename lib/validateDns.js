const debug = require('debug');
const { Resolver } = require('node:dns').promises;

const dlog = debug('brettski:ipwatch:ipvalidate');
const resolver = new Resolver();

exports.validateAtCloudFlare = (hostname, expectedIp) => {
  dlog('validateAtCloudFlare %s is %s', hostname, expectedIp);
  resolver.setServers(['1.1.1.1', '1.0.0.1']);
  return resolver
    .resolve4(hostname)
    .then(addresses => {
      dlog('addresses: %o', addresses);
      return {
        isEqual: addresses[0] === expectedIp,
        message: `returned address: ${addresses}`,
      };
    })
    .catch(err => ({
      isEqual: false,
      message: err.message,
    }));
};
