# sweep

Delete expired items in [Pocket](https://getpocket.com)

## Installation

```
$ go get github.com/naoty/sweep
```

## Usage

```
$ sweep
```

## Configuration

```
POCKET_CONSUMER_KEY: A consumer key for your Pocket application
POCKET_ACCESS_TOKEN: An access token for your Pocket application
POCKET_EXPIRATION: An expiration hours after which items are deleted (Default: 24)
```

## sweeper

You can deploy a job to run sweep on Heroku.

1. Get a consumer key and an access token of your Pocket application.
2. Deploy sweeper from below Heroku button.
3. Add a new job at Heroku scheduler dashboad.

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Author

[naoty](https://github.com/naoty)

