import os
import json
from dotenv import load_dotenv

def configMissing(configSetting):
  raise(ValueError("Missing required config setting: " + configSetting))

def load_config():
  load_dotenv()
  config = {
    'end_point': os.getenv("ENDPOINT_CHK") or configMissing("ENDPOINT_CHK"),
    'test_location': os.getenv("TEST_LOCATION") or 'Test Location',
    'update_on_change': json.loads(os.getenv('UPDATE_ON_CHANGE')) or False,
    'postmark': {
      'token': os.getenv("POSTMARK_TOKEN") or configMissing("POSTMARK_TOKEN"),
      'email_to': os.getenv("EMAIL_TO") or configMissing("EMAIL_TO"),
      'email_from': os.getenv("EMAIL_FROM") or configMissing("EMAIL_FROM"), 
    }
  }

  # print(f'have config: {config}')
  return config
  # if not os.getenv("POSTMARK_TO_EMAIL"):
  #   configMissing("POSTMARK_TO_EMAIL")
