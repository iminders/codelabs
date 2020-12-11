install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/rakyll/statik \
		github.com/bufbuild/buf/cmd/buf

gen_proto:
	protoc -I . -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out ./gen/go/ --go_opt paths=source_relative \
	   --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative \
	   proto/echo_service/echo_service.proto

gw:
	protoc -I . -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out ./gen/go \
     --grpc-gateway_opt logtostderr=true \
     --grpc-gateway_opt paths=source_relative \
     --grpc-gateway_opt generate_unbound_methods=true \
     proto/echo_service/echo_service.proto

openapi:
	protoc -I . -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true proto/echo_service/echo_service.proto

protos:
	make gen_proto
	make gw
