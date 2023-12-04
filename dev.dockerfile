FROM golang:1.20-alpine

RUN apk update && apk upgrade
RUN apk add --no-cache sqlite sqlite-libs build-base
RUN sqlite3 --version

WORKDIR /app

COPY go.mod .

RUN go mod download

RUN go install github.com/mitranim/gow@latest

COPY . .

CMD ["gow", "run", "."]