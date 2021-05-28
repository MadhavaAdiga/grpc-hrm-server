generate:
	protoc -I protos/ protos/*.proto --go_out=protos/
	protoc -I protos/ protos/*.proto --go-grpc_out=protos/