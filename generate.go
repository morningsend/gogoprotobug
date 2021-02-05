package main

//go:generate protoc --proto_path=$GOPATH/src:$PWD/vendor:$PWD/vendor/github.com/gogo/protobuf/protobuf:. --gogofaster_out=.,plugins=grpc:. pb/mystruct.proto
