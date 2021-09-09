package main

import (
	"github.com/dandyhuang/cmd_tools/cmd"
	_ "github.com/jhump/protoreflect/desc"
	_ "github.com/jhump/protoreflect/desc/protoparse"
	_ "google.golang.org/grpc/metadata"
	_ "google.golang.org/grpc/status"
)

func main() {
	cmd.Execute()
}
