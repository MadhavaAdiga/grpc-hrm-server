generate:
	protoc -I protos/ protos/*.proto --go_out=protos/hrm
	protoc -I protos/ protos/*.proto --go-grpc_out=protos/hrm