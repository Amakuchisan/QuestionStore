# QuestionBox

This is a bulletin board format Web application. Users post questions and comments on questions.

[![CircleCI](https://circleci.com/gh/Amakuchisan/QuestionBox/tree/master.svg?style=svg)](https://circleci.com/gh/Amakuchisan/QuestionBox/tree/master)

## Overview

- Login with OAuth2 with Google
- Post questions
- Post comments on questions

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
$ ./bin/tts
```

## Launch local environment

```
$ docker-compose up -d
```

**Execute SQL**
```
$ docker-compose exec db mysql -h localhost -u tts -p tts
```

## Future work

- Use Nginx