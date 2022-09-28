FROM golang
COPY . /go/src/client
WORKDIR /go/src/client
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 8080
