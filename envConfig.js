function configMissing(configSetting) {
  throw new Error(`missing required .env setting: ${configSetting}`);
}

const envConfig = () => ({
  endpoint: process.env.ENDPOINT_CHK || configMissing('ENDPOINT_CHK'),
  testLocation: process.env.TEST_LOCATION || 'Test Location',
  postmark: {
    token: process.env.POSTMARK_TOKEN || POSTMARK_TOKEN,
    emailto: process.env.EMAIL_TO || configMissing('EMAIL_TO'),
    emailfrom: process.env.EMAIL_FROM || configMissing('EMAIL_FROM'),
  },
  CloudFlare: {
    apiToken: process.env.CLOUDFLARE_API_TOKEN || null,
    endpoint: 'https://api.cloudflare.com/client/v4/',
  }
})

module.exports = envConfig();
