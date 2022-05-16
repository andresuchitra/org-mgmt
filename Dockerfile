# Builder
FROM golang:1.17.10-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .
COPY .env.docker /app/.env

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app
COPY .env.docker /app/.env
EXPOSE 9090

COPY --from=builder /app/engine /app

CMD /app/engine