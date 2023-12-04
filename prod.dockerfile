FROM golang:1.20-alpine as builder

RUN apk update && apk upgrade
RUN apk add --no-cache sqlite sqlite-libs build-base
RUN sqlite3 --version

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

FROM alpine:3.18.4

RUN apk update && apk upgrade
RUN apk add --no-cache sqlite sqlite-libs
RUN sqlite3 --version


WORKDIR /app

COPY --from=builder /app/eco-challenge .

RUN mkdir upload && cd upload && mkdir common
RUN mkdir data