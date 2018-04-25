FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get -v -u ./...

RUN go clean
RUN GOOS=linux GOARCH=amd64 go build

CMD ["link"]
