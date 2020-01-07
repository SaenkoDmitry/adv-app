FROM golang:latest

WORKDIR $GOPATH/src/advertisement-app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
