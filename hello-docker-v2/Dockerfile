FROM golang:1.16-alpine AS builder
WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o /hello-docker

FROM alpine:3.14
COPY --from=builder /hello-docker /
CMD ["/hello-docker"]
