# for sending notication emails
from postmarker.core import PostmarkClient
import envConfig

env_config = envConfig.load_config()
postmark = PostmarkClient(server_token=env_config['postmark']['token'])

def send(subject, body):
  if len(subject) < 1:
    subject = '[action required] IP Address notification. Somthing''s different'
  
  postmark.emails.send(
    From=env_config['postmark']['email_from'],
    To=env_config['postmark']['email_to'],
    Subject=subject,
    TextBody=body,
  )