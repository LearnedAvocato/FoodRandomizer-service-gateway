# docker image just to build
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

ENV  GO111MODULE=on
ENV  CGO_ENABLED=0

RUN go build -o broker-service ./cmd/api

RUN chmod +x /app/broker-service

# small docker image to run 
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/broker-service /app

CMD ["./app/broker-service"]
EXPOSE 80