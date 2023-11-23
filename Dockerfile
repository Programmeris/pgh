FROM golang:alpine as builder
WORKDIR /build
COPY . .
RUN go build -o pgh main.go
CMD [". /pgh"]
