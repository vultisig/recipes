package main

//go:generate protoc --proto_path=/usr/local/include/include -I/usr/local/include/include -I./proto --go_opt=paths=source_relative --go_out=./types proto/constraint.proto proto/rule.proto proto/parameter_constraint.proto proto/resource.proto proto/policy.proto proto/recipe_specification.proto proto/scheduling.proto

func main() {

}
