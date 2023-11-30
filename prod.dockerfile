FROM golang as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /build/api_file

FROM alpine:3.18.4

RUN apk update

WORKDIR /app

COPY --from-builder /app/build/api_file ./app

RUN mkdir upload
RUN mkdir data