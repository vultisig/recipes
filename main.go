package main

//go:generate protoc -I./proto --go_opt=paths=source_relative --go_out=./types proto/policy.proto proto/rule.proto proto/constraint.proto proto/resource.proto
func main() {

}
