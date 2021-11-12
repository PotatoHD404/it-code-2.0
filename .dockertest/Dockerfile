FROM golang:1.17.2-alpine3.14 AS builder
WORKDIR /build
COPY ./cart/ /build/
RUN go build -o cart

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /build/cart /app/cart
CMD [ "/app/cart" ]
