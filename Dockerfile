##  Building Stage
FROM golang:1.13.0-alpine as builder

RUN apk add --no-cache make git

WORKDIR /src
COPY go.mod /src/go.mod
RUN go mod download
COPY . .
RUN make build
RUN go get gopkg.in/urfave/cli.v2@master
RUN go get github.com/oxequa/realize


## Running Stage

FROM alpine:3.10.2
WORKDIR /app
COPY --from=builder /src/bin/tts /app/tts
# Template files
COPY --from=builder /src/templates/ /app/templates
# Static files (js|css|img)
COPY --from=builder /src/static/ /app/static

CMD /app/tts
