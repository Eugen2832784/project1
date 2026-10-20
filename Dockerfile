FROM golang:1.20-alpine AS builder

RUN apk --no-cache add bash

WORKDIR /usr/local/src

COPY ["main.go", "./"]

RUN go build -o app main.go

CMD ["./app"]