#/bin/bash
protoc -I protoc --go_out=plugins=grpc:protoc \
	basic.proto	\
