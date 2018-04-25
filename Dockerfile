FROM golang:latest

WORKDIR /go/src/github.com/xiaosongfu/link

COPY . .

# install and run dep
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# build
RUN go clean
RUN go install

CMD ["link"]
