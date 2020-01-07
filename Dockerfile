FROM golang:latest

ARG db_host
ENV db_host "$db_host"

WORKDIR $GOPATH/src/advertisement-app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
