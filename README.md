# Golang Bindings for the Telegram Bot API

# Headline

This is a maintained fork of github.com/webnice/tba that preserves the original design while staying current with the latest Telegram Bot API specifications.

## Features

- Complete coverage of the Telegram Bot API
- Simple and intuitive interface
- Minimal dependencies
- Production-ready and actively maintained
- Preserves original API structure for easy migration

## Quick Start

Here's a simple echo bot using polling:

```go
package main

import (
    "log"
    "time"
    api "github.com/webnice/tba/v9"
)

func main() {
    bot, err := api.NewBotAPI("BOT_TOKEN")
    if err != nil {
        panic(err)
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    updateConfig := api.NewUpdate(0)
    updateConfig.Timeout = 60
    updatesChannel := bot.GetUpdatesChan(updateConfig)

    // Optional: Clear initial updates
    time.Sleep(time.Millisecond * 500)
    updatesChannel.Clear()

    for update := range updatesChannel {
        if update.Message == nil {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        msg := api.NewMessage(update.Message.Chat.ID, update.Message.Text)
        msg.ReplyParameters.MessageID = update.Message.MessageID

        bot.Send(msg)
    }
}
```

## Webhook Setup

For webhook implementation:
1. Check the [examples](./examples) folder for examples of webhook implementations
2. Use HTTPS (required by Telegram)
   - Generate SSL certificate:
    ```bash
    openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes
    ```
   - Bring your own SSL certificate. [Let's Encrypt](https://letsencrypt.org) is recommended for free TLS certificates in production.

## Contributing

- Issues and feature requests are welcome
- Pull requests should maintain the existing design philosophy
- The focus is on providing a clean API wrapper without additional features
