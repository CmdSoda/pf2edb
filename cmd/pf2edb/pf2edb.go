package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		Foo string
		Bar bool
	}
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar)
}
