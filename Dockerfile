FROM golang:1.12-stretch

WORKDIR /go/src/app

COPY main.go .

RUN go build -v

ENTRYPOINT [ "./app" ]
