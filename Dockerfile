# Build stage
FROM golang:1.17-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .

# Env
ENV HOST_ADDRESS "0.0.0.0"
ENV HTTP_PORT "5000"

EXPOSE 5000
CMD ["/app/main"]