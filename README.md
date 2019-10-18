# QuestionStore

This is a bulletin board format Web application. Users post questions and comments on questions.

[![CircleCI](https://circleci.com/gh/Amakuchisan/QuestionStore/tree/master.svg?style=svg&circle-token=684da293d09fdd7e3ec93f8363c827f70ffae5b7)](https://circleci.com/gh/Amakuchisan/QuestionStore/tree/master)

## Overview

- Login with OAuth2 with Google
- Post questions
- Reply to questions

## Dependencies

- Go: 1.13.0
- Echo: 3.3.10
- MySQL: 8.0.17
- Docker: 19.03.2
- Docker Compose: 1.24.1

## Building

```
$ make build
```

## Running

```
$ ./bin/qs
```

## Launch local environment

```
$ docker-compose up -d
```

## Future work

- Email Notification
- Export static file (Nginx delivery)
- Testing
- UI improvements
