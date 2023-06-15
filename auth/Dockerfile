FROM golang:1.19-alpine as builder

COPY . /auth

WORKDIR /auth

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o auth-service ./cmd/main.go

#**# Это контейнер с ранее собранным приложением**

FROM alpine:latest as server

WORKDIR /root/

COPY --from=builder /auth/auth-service .

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

RUN chown root:root auth-service

EXPOSE 8080

CMD ["./auth-service"]