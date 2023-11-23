FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go build -o pgh main.go

FROM alpine as app

COPY --from=builder /build/pgh .

CMD [". /pgh"]
