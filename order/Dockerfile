FROM golang:1.19-alpine as builder

COPY . /order

WORKDIR /order

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o order-service ./cmd/main.go

#**# Это контейнер с ранее собранным приложением**

FROM alpine:latest as server

WORKDIR /root/

COPY --from=builder /order/order-service .

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

RUN chown root:root order-service

EXPOSE 8004

CMD ["./order-service"]