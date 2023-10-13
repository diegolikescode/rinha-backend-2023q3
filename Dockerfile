FROM golang:1.21

# RUN apt-get update && apt-get install -y netcat-traditional

WORKDIR /app

# COPY wait-for-postgres.sh /usr/local/bin

# RUN chmod +x /usr/local/bin/wait-for-postgres.sh

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 9999

# CMD ["wait-for-postgres.sh", "db", "./main"]
CMD ["./main"]

