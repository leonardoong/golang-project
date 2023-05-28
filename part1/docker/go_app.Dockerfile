FROM golang:1.20

WORKDIR /go/src/app

COPY . .

RUN make build-grpc

CMD [ "./cmd/grpc/grpc" ]