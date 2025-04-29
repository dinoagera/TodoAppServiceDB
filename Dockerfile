FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/api-db/main.go

FROM alpine:3.18

COPY --from=builder /app/main /main

COPY migrations /migrations

EXPOSE 8081

CMD ["/main"]