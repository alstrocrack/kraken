FROM golang:1.20-alpine

RUN mkdir /echo
COPY cmd/kraken/main.go /echo

CMD ["go", "run", "/echo/main.go"]
