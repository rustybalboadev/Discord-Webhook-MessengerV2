# Discord Webhook Messenger

# Setup
* Open Config.json
* Input Webhook URL
* Avatar URL & Webhook Username are **optional**

# Usage
```
usage: Webhook Sender [-h|--help] -m|--message "<value>" [-t|--timeout
                      <integer>] [-a|--amount <integer>]

                      Spam a webhook with a message

Arguments:

  -h  --help     Print help information
  -m  --message  Message to Send Through Webhook
  -t  --timeout  Time Between Messages In Seconds
  -a  --amount   Amount Of Times To Send Message
```

# Example
``go run webhook.go -m hello -t 5 -a 50`` **Send the message "hello" every 5 seconds 50 times**

**MADE BY cookie#0003 ON DISCORD**