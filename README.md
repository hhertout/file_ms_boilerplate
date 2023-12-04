# Presentation

## Getting started

## Dev

### config
```yaml
services:
  file:
    build:
      context: .
      dockerfile: dev.dockerfile
    env_file:
      - ./.env
    ports:
      - 8080:8080
    volumes:
      - .:/app
      - db:/app/data
      - upload:/app/upload
volumes:
  db:
  upload:
```

## Production

### config
```yaml
services:
  file:
    build:
      context: .
      dockerfile: prod.dockerfile
    environment:
      - ENV=prod
    env_file:
      - ./.env
    ports:
      - 8080:8080
    volumes:
      - db:/app/data
      - upload:/app/upload
    command: "./eco-challenge"
volumes:
  db:
  upload:
```