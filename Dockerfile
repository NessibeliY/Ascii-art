FROM golang:1.20.1

WORKDIR /app

COPY . .

RUN go build -o ascii-art-web ./cmd/web/

CMD ["./ascii-art-web"]