FROM golang
COPY . /go/src/books
WORKDIR /go/src/books
RUN go get .
ENTRYPOINT go run main.go
EXPOSE 5000
