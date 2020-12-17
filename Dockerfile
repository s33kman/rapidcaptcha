FROM golang:1.14

WORKDIR /go/src/chat
COPY .. .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["chat"]