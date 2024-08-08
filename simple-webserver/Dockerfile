FROM golang:alpine

WORKDIR /simpleserver

COPY . /simpleserver

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]