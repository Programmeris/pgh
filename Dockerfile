FROM golang:1.17-alpine as builder

WORKDIR /build

COPY . .

RUN go build -o pgh main.go

FROM alpine as app

COPY --from=builder /build/pgh .

CMD [". /pgh"]
