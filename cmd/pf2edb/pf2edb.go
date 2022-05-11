package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
)

// usage: pf2edb.exe <dbfile>
func main() {
	var args struct {
		Foo string
		Bar bool
	}
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar)
}
