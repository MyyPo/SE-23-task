# App for SE-23 test task

## Features and notes

- Includes RESTful JSON API
- Uses simple fs db implementation that stores records in files by the record's first letter, records are separated by newlines
- Uses free [Coingecko API](https://www.coingecko.com/en/api), according to the docs they have a limit of 10-30 requests/min, so it shoudn't throttle... theoretically
- Uses [Gin](https://github.com/gin-gonic/gin)
- I used [Mailtrap](https://mailtrap.io) for testing it

## How to start

- Fill .env file according to .env.example
- Use docker compose **docker compose up -d --build**

## Endpoints

- *GET* -> **/rate** endpoint returns formatted current BTC to UAH exchange rate
- *POST* -> **/subscribe** endpoint subscribes a new email (stores it in a simple filesystem database), expects "email" JSON field filled
- *POST* -> **/sendEmails** endpoint sends emails with the currect BTC to UAH exchange rate to subscribed emails
