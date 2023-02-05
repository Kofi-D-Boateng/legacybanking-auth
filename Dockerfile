FROM golang:1.18.3-alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main

CMD ["/app/main"]