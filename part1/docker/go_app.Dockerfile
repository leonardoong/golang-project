FROM golang:1.20

WORKDIR /go/src/app

COPY . .

RUN go mod tidy
RUN go mod vendor

RUN make build-grpc

CMD [ "./cmd/grpc/grpc" ]