FROM golang:1.14.2-alpine
RUN apk add git ca-certificates
WORKDIR /go/src/app
COPY . .
RUN go get -v
RUN go build -v
CMD ["./app"]