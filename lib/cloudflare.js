const fetch = require('node-fetch');
const debug = require('debug');
const envConfig = require('../envConfig');

const dlog = debug('brettski:ipwatch:cloudflare');

const status = {
  isOk: false,
  message: '',
};

const cloudFlare = () => {
  dlog('cloudFlare instance created');

  const { apiToken } = envConfig.CloudFlare;
  let { apiBase } = envConfig.CloudFlare;
  if (!apiToken || !apiBase) {
    throw new Error(
      'Missing api token or api base url for CloudFlare, unable to proceed',
    );
  }
  if (apiBase.endsWith('/')) {
    apiBase = apiBase.substring(0, apiBase.length - 1);
  }
  const headers = {
    Authorization: `Bearer ${apiToken}`,
    'Content-Type': 'application/json',
  };
  const options = {
    headers,
    method: 'GET',
  };

  status.isOk = true;
  status.message = 'instance created';

  function getZoneId(zoneDomain) {
    const url = `${apiBase}/zones?name=${zoneDomain}`;
    dlog('getZoneId url: %s', url);
    return fetch(url, options)
      .then(res => res.json())
      .then(({ result }) => {
        // dlog('json: %O', result);
        if (!result || result?.lenth < 1) {
          return null;
        }
        const [zone] = result;
        return {
          id: zone.id,
          name: zone.name,
        };
      });
  }

  function getDnsRecordId(zoneId, dnsDomainName) {
    const url = `${apiBase}/zones/${zoneId}/dns_records?name=${dnsDomainName}`;
    dlog('getDnsRecord url: $s', url);
    return fetch(url, options)
      .then(res => res.json())
      .then(({ result }) => {
        dlog('json %O', result);
        if (!result || result?.lenth < 1) {
          return null;
        }
        const [dnsRecord] = result;
        return {
          id: dnsRecord.id,
          zoneId: dnsRecord.zone_id,
          name: dnsRecord.name,
          zoneName: dnsRecord.zone_name,
          type: dnsRecord.type,
          content: dnsRecord.content,
        };
      });
  }

  function updateDnsRecordIP({ zoneId, dnsRecordId, newIp }) {
    dlog('updating zone %s, record %s, with ip %s', zoneId, dnsRecordId, newIp);
    const url = `${apiBase}/zones${zoneId}/dns_records/${dnsRecordId}`;
    dlog('updateDnsRecordIp url: %s', url);
    const _options = {
      headers,
      method: 'PATCH',
      body: JSON.stringify({ content: newIp }),
    };
    return fetch(url, _options);
  }

  return { status, getZoneId, getDnsRecordId, updateDnsRecordIP };
};

module.exports = cloudFlare;
