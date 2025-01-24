# RSS feed aggregator

An [RSS](https://en.wikipedia.org/wiki/RSS) feed aggregator in Go! It's a web server that allows clients to:

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

# DB Migration Using GOOSE

## Installation

```
    go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Running UP Migration

```
    goose up -env=goose.en
```

# ORM Using SQLC

## Installation

```
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Code Generation

```
    sqlc generate
```
