FROM golang:1.14.0

WORKDIR /app
ADD . /app

ENV GO111MODULE="on"
RUN go mod download

CMD ["go", "run", "server.go"]