FROM golang:alpine
WORKDIR /build
COPY . .
RUN go build -o pgh main.go
CMD [". /pgh"]
