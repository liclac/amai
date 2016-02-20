FROM golang:1.6
RUN mkdir -p /go/src/github.com/uppfinnarn/amai
WORKDIR /go/src/github.com/uppfinnarn/amai
ADD . /go/src/github.com/uppfinnarn/amai
RUN go get ./...
RUN go install
