FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY . .

RUN go build -o kraken ./cmd/kraken

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/kraken .

CMD ["./kraken"]
