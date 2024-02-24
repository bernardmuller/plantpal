FROM golang:1.22.0-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /domain-app

EXPOSE 8080

CMD ["/domain-app"]
