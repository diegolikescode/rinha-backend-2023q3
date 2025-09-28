FROM golang:1.25-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./src ./src

RUN go build -o rinha ./cmd/rinha.go

CMD ["./rinha"]
