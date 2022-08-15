FROM golang:1.16
WORKDIR /app
ADD . .
RUN go mod download
RUN go build -o /hello-docker
CMD ["/hello-docker"]
