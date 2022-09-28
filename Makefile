gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server --experimental_allow_proto3_optional
	protoc --proto_path=proto proto/*.proto --go_out=api --go-grpc_out=api --experimental_allow_proto3_optional

clean:
	rm -rf server/pb/
	rm -rf api/pb/

server:
	go run server/main.go

api:
	go run server/main.go

build-api:
	cd api && docker build . -t hutchery/api-service:v1.0

push-api:
	docker push hutchery/api-service:v1.0

build-server:
	cd server && docker build . -t hutchery/books-service:v1.0

push-server:
	docker push hutchery/books-service:v1.0
