# Telegram Shell Bot

Run shell command & take screenshots of your PC using this bot.

## How to Run
- First create a Telegram bot using [BotFather]
- Then run `go get github.com/faem/telegram-sh-bot`
- Finally, `telegram-sh-bot --token <bot-api-token> --admin <telegram-username-of-admin>`
- Enjoy!

## Run from Docker
- `docker run fahimabrar/bot --token <bot-api-token> --admin <telegram-username-of-admin>`

## Available Flags

```
  -admin string
        Username of the admin (default "fahim_abrar")
  -debug
        Print error info to debug
  -poll_time int
        Response time of bot (default 100)
  -token string
        Token of your bot
```
## Bot Commands

- `/hello`: simply greets the user (anyone can run it)
- `/getss`: takes screenshots of the pc and sends it to the user (only admin can run it)
- `/sh <valid-shell-command>`: runs the command where the bot is running and sends any output returned by the command or error (only admin can run it)
 
 ## Bot ScreenShots
Hello (`/hello`)              |      Get Screenshots (`/getss`)                |        Shell (`/sh cat Dockerfile`) 
:-------------------------:|:-------------------------:|:-------------------------:
![img](https://i.ibb.co/Wn6Rm4G/ss2.jpg) | ![img](https://i.ibb.co/g4BMSSF/ss3.jpg) | ![img](https://i.ibb.co/C2JWdfr/ss1.jpg) 

[BotFather]: https://core.telegram.org/bots#6-botfather
