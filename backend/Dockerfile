FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /server cmd/server/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /server /usr/local/bin/server

WORKDIR /data

EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/server"]
