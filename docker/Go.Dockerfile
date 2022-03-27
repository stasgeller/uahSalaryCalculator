# syntax=docker/dockerfile:1
FROM golang:1.16-alpine as dev

WORKDIR /uahSalaryBot

ARG BOT_API_KEY=${TG_BOT_API_KEY}

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download

RUN go install github.com/go-delve/delve/cmd/dlv@latestY
RUN go get -d -v ./...

RUN go build -o uahSalaryBot ./cmd

# DEBUG MODE
FROM dev as debug

EXPOSE 8080 40000
#
##WORKDIR /
COPY --from=dev /go/bin/dlv /
COPY --from=dev . .
##
#CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./cmd"]

CMD ["./uahSalaryBot"]