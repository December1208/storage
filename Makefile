API_PROTO_FILES=$(shell find api -name *.proto)


# generate grpc code
.PHONY: grpc
grpc:
	protoc --proto_path=. \
           --proto_path=third_party \
           --go_out=paths=source_relative:. \
           $(API_PROTO_FILES)
