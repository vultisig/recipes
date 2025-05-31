package main

//go:generate protoc -I./proto --go_opt=paths=source_relative --go_out=./types proto/constraint.proto proto/rule.proto proto/parameter_constraint.proto proto/resource.proto proto/policy.proto proto/recipe_specification.proto

func main() {

}
