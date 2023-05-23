# build environment
FROM golang:1.19-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum files first
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the project files
COPY . .

#ENV CGO_ENABLED=0
RUN go build -o /app/gateway-service ./cmd/gateway

# production environment
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gateway-service .

CMD ["./gateway-service"]
EXPOSE 80