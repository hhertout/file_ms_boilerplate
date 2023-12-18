FROM golang:1.20-alpine

RUN apk update && apk upgrade
RUN apk add  \
    sqlite  \
    sqlite-libs  \
    build-base  \
    make

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . .

RUN mkdir data
RUN mkdir upload && cd upload && mkdir common

RUN go mod download

CMD ["make", "watch"]