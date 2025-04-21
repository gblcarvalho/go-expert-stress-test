FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o stress-test ./cmd/cli/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stress-test .

ENTRYPOINT ["./stress-test"]
