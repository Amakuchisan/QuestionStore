##  Building Stage
FROM golang:1.13.0-alpine as builder

RUN apk add --no-cache make git \
  && go get github.com/oxequa/realize


WORKDIR /go/src/github.com/Amakuchisan/QuestionBox
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN make build


## Running Stage

FROM alpine:3.10.2
WORKDIR /app
COPY --from=builder /go/src/github.com/Amakuchisan/QuestionBox/bin/tts /app/tts
# Template files
COPY --from=builder /go/src/github.com/Amakuchisan/QuestionBox/templates/ /app/templates
# Static files (js|css|img)
COPY --from=builder /go/src/github.com/Amakuchisan/QuestionBox/static/ /app/static

CMD /app/tts
