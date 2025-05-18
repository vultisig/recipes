package main

//protoc -I/usr/local/include -I$HOME/dev/thebitpress/bitpress/proto/ -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,paths=source_relative:$HOME/dev/thebitpress/bitpress/.  $HOME/dev/thebitpress/bitpress/proto/types/*"

//go:generate protoc -I./proto --go_opt=paths=source_relative --go_out=./types proto/policy.proto proto/rule.proto proto/constraint.proto proto/resource.proto
func main() {

}
