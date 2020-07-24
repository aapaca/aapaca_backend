FROM golang:1.14.0

RUN go get "github.com/go-sql-driver/mysql"
RUN go get -u "github.com/labstack/echo/..."

WORKDIR /app
ADD . /app

CMD ["go", "run", "server.go"]