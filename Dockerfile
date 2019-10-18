##  Building Stage
FROM golang:1.13.0-alpine as builder

ENV LANG en_US.UTF-8

RUN apk add --update --no-cache make git tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone && \
    apk del tzdata && \
    go get github.com/oxequa/realize

WORKDIR /go/src/github.com/Amakuchisan/QuestionStore
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN make build


## Running Stage

FROM alpine:3.10.2

ENV LANG en_US.UTF-8
RUN apk add --update --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone && \
    apk del tzdata

WORKDIR /app

COPY --from=builder /go/src/github.com/Amakuchisan/QuestionStore/bin/qs /app/qs
# Template files
COPY --from=builder /go/src/github.com/Amakuchisan/QuestionStore/templates/ /app/templates
# Static files (js|css|img)
COPY --from=builder /go/src/github.com/Amakuchisan/QuestionStore/static/ /app/static

CMD /app/qs
