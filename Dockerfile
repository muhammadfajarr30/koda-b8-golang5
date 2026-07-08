FROM golang:alpine AS builder

WORKDIR /app

COPY . . 

RUN go mod tidy

RUN go build -o minitask-authentication

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/minitask-authentication /app/

CMD ["/app/minitask-authentication"]