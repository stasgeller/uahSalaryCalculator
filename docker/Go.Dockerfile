# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /uahSalaryBot

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download

RUN go build -o uahSalaryBot ./cmd

EXPOSE 8080

CMD ["./uahSalaryBot"]
