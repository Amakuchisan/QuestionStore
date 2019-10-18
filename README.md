# QuestionBox

This is a bulletin board format Web application. Users post questions and comments on questions.

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

## Use Docker Compose

**Build**
```
$ docker-compose build
```

**Launch**
```
$ docker-compose up -d
```

**Down**
```
$ docker-compose down
```

**Logging**
```
$ docker-compose logs -f
```

**Execute SQL**
```
$ docker-compose exec db mysql -h localhost -u tts -p tts
```

[![CircleCI](https://circleci.com/gh/Amakuchisan/QuestionBox/tree/master.svg?style=svg)](https://circleci.com/gh/Amakuchisan/QuestionBox/tree/master)
