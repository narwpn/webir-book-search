FROM golang:bookworm AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o crawler-executable

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/crawler-executable /app/crawler-executable

RUN chmod +x /app/crawler-executable

ENTRYPOINT [ "/app/crawler-executable" ]