# Presentation

## Getting started

## Dev

### config
```yaml
services:
  file:
    build:
      context: .
      dockerfile: Dockerfile
    env_file:
      - ./.env
    ports:
      - 4568:4568
    volumes:
      - .:/app
```

## Production

### Dockerfile

```dockerfile
FROM golang:1.20-alpine as builder

RUN apk update && apk upgrade
RUN apk add --no-cache sqlite sqlite-libs build-base

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

FROM alpine:3.18.4

RUN addgroup -g 1000 golang \
    && adduser -u 1000 -G golang -s /bin/sh -D golang

RUN apk update && apk upgrade
RUN apk add --no-cache sqlite sqlite-libs
RUN sqlite3 --version

WORKDIR /app

COPY --from=builder /app/eco-challenge .

RUN mkdir upload && cd upload && mkdir common
RUN mkdir data
```

### config
```yaml
services:
  file:
    build:
      context: .
      dockerfile: Dockerfile
    env_file:
      - ./.env
    ports:
      - 4568:4568
    volumes:
      - db:/app/data
      - upload:/app/upload
    command: "./eco-challenge"
volumes:
  db:
  upload:
```