FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -vi ./...
RUN go install -v ./...
 

CMD go run ./cmd/transcriber.go
