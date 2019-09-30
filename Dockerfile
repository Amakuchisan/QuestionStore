##  Building Stage
FROM golang:1.13.0-alpine as builder

RUN apk add --no-cache make

WORKDIR /src
COPY go.mod /src/go.mod
RUN go mod download
COPY . .
RUN make build

## Running Stage

FROM alpine:3.10.2
WORKDIR /app
COPY --from=builder /src/bin/tts /app/tts
# Template files
COPY --from=builder /src/templates/ /app/templates
# Static files (js|css|img)
COPY --from=builder /src/static/ /app/static
EXPOSE 8080

CMD /app/tts
