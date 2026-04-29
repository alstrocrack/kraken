FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o kraken ./cmd/kraken

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/kraken .

EXPOSE 8080

CMD ["./kraken"]