# Builder
FROM golang:1.17.10-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY go.* .
RUN go mod download
COPY . .
COPY .env.docker /app/.env

RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o engine main.go

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