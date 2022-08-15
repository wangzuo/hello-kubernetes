FROM golang:1.19 AS builder
WORKDIR /go/src/hello-docker
ADD . .
RUN go build -o main .

FROM alpine
WORKDIR /root
COPY --from=builder /go/src/hello-docker/main .
CMD ["./main"]
