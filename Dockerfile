FROM golang:1.14

WORKDIR /go/src/rapidcaptcha
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["rapidcaptcha"]