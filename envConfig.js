function configMissing(configSetting) {
  throw new Error(`missing required .env setting: ${configSetting}`);
}

const envConfig = () => ({
  endpoint: process.env.ENDPOINT_CHK || configMissing('ENDPOINT_CHK'),
  testLocation: process.env.TEST_LOCATION || 'Test Location',
  updateOnChange: JSON.parse(process.env.UPDATE_ON_CHANGE ?? false),
  postmark: {
    token: process.env.POSTMARK_TOKEN || configMissing('POSTMARK_TOKEN'),
    emailto: process.env.EMAIL_TO || configMissing('EMAIL_TO'),
    emailfrom: process.env.EMAIL_FROM || configMissing('EMAIL_FROM'),
  },
  CloudFlare: {
    apiToken: process.env.CF_BEARER_TOKEN ?? null,
    apiBase: 'https://api.cloudflare.com/client/v4/',
    zoneDomain: process.env.CF_ZONE_DOMAIN ?? null,
    recordDomain: process.env.CF_DNS_RECORD_DOMAIN ?? null,
  },
});

module.exports = envConfig();
