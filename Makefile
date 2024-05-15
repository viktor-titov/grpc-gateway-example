run:
	go run internal/services/grpc/greeting/greeting.go

buf-check:
	buf lint proto/grpc-gateway-example 

buf-gen: 
	buf generate proto/grpc-gateway-example --template proto/grpc-gateway-example/gen.server.yaml
	buf generate proto/grpc-gateway-example --template proto/grpc-gateway-example/gen.swagger.yaml

buf-gen-client:
	buf generate proto/grpc-gateway-example --template proto/grpc-gateway-example/gen.client.yaml