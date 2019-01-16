rtcamp
===

This package allows you to search retweeted users of the specified tweet and write them into a csv file.

You need `.env` file as below to run this script:

```
CONSUMER_API_KEY="twitter consumer api key"
CONSUMER_API_SECRET="twitter consumer api secret"
ACCESS_TOKEN="twitter access token"
ACCESS_TOKEN_SECRET="twitter access token secret"
```

This package uses [Modules](https://github.com/golang/go/wiki/Modules), so you can run this package by `GO111MODULE=on go run main.go`.