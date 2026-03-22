FROM golang:1.23.1 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/server/main.go

FROM debian:12-slim
COPY --from=builder /app/main /main
CMD ["/main"]