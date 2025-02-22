FROM golang:1.22.6-alpine AS builder

WORKDIR /app

COPY scripts/wait-for-it.sh /usr/bin/wait-for-it
RUN chmod +x /usr/bin/wait-for-it

CMD ["wait-for-it", "postgres:5432", "--", "./main"]

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM alpine:3.17

WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080
CMD ["./main"]

