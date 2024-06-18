FROM golang:1.22.0-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd services/web && CGO_ENABLED=0 GOOS=linux go build -o /web-service

EXPOSE 8000

CMD ["/web-service"]
