# docker image just to build
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

ENV  GO111MODULE=on
ENV  CGO_ENABLED=0

RUN go build -o gateway-service ./cmd/api

RUN chmod +x /app/gateway-service

# small docker image to run 
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/gateway-service /app

CMD ["./app/gateway-service"]
EXPOSE 80