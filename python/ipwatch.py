import requests
from tinydb import TinyDB, Query
from tinydb.operations import increment as dbinc, set as dbset
import datetime
import envConfig

env_config = envConfig.load_config()
try:
  result = requests.get(env_config['end_point'])

except Exception as err:
  # TODO: log this
  print(f'logging {err}')

if result.status_code == 200:
  jresult = result.json();
  cur_ip = jresult['x-forwarded-for']
  if len(cur_ip) > 0:
    db = TinyDB('./data.json')
    Log = Query()
    now = datetime.datetime.now(datetime.UTC).isoformat()
    row = db.get(Log.ip == cur_ip)
    if row is not None:
      # couldn't figure out how to do mulple operations on one doc
      # in a single call, e.g. using `update_multiple()`
      db.update(dbinc('count'), doc_ids=[row.doc_id])
      db.update(dbset('timestamp', now), doc_ids=[row.doc_id])
      print(db.all())
    else:
      # new ip
      pass

  else:
    # send error email (no ip in x-forwarded-for)
    pass

else:
  # record issue
  pass

print('done', result)

