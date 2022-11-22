FROM golang:1.19-alpine
WORKDIR /app
COPY ../Go/go.mod ./
COPY ../Go/go.sum ./
RUN go mod download
COPY ../Go/icmp-probe-service.go ./
RUN go build -o ./docker-probe-service

CMD ["/app/docker-probe-service"]
