FROM golang:latest
ADD . /go/src/github.com/slovnik/slovnik_api
WORKDIR /go/src/github.com/slovnik/slovnik_api
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go install github.com/slovnik/slovnik_api
ENTRYPOINT /go/bin/slovnik_api
EXPOSE 8080
