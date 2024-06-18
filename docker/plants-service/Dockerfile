FROM golang:1.22.0-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cd services/plants-service && CGO_ENABLED=0 GOOS=linux go build -o /plants-service

EXPOSE 8001
EXPOSE 9001

CMD ["/plants-service"]
