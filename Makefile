gen:
	protoc --proto_path=common/proto common/proto/*.proto  --go_out=common/proto/pb --go-grpc_out=common/proto/pb --grpc-gateway_out=common/proto/pb
