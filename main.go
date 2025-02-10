package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/protoc-gen/protoc-gen-go-errors/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	//go:embed VERSION
	Version string
)

var showVersion = flag.Bool("version", false, "print the version and exit")

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-errors %v\n", Version)
		return
	}
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(g *protogen.Plugin) error {
		g.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range g.Files {
			if !f.Generate {
				continue
			}
			gen.GenerateFile(g, f)
		}
		return nil
	})
}
