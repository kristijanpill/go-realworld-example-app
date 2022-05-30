module github.com/kristijanpill/go-realworld-example-app/api_gateway

go 1.18

replace github.com/kristijanpill/go-realworld-example-app/common => ../common

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.2
	github.com/kristijanpill/go-realworld-example-app/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.46.2
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220526153639-5463443f8c37 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220527130721-00d5c0f3be58 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
