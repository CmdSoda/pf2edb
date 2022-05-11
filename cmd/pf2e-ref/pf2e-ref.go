package main

import (
	"fmt"
	"github.com/CmdSoda/pf2edb/internal/pf2e"
	"github.com/alexflint/go-arg"
	"log"
	"os"
)

// usage: pf2e-ref.exe [-c type] <search-string>
func main() {
	var args struct {
		Type   string `arg:"-t"`
		Search string `arg:"positional, required"`
	}
	arg.MustParse(&args)

	ctx := pf2e.NewContext()
	err := ctx.Config.Load("./config.json")
	if err != nil {
		panic("verdammt")
	}

	systemPath := ctx.Config.DataFolder + "\\systems\\" + ctx.Config.SystemName

	if _, err := os.Stat(systemPath); os.IsNotExist(err) {
		log.Fatal(systemPath + " does not exist")
		return
	}
	systemFilename := systemPath + "\\system.json"
	sys, errLoadSystem := pf2e.NewSystemFromFilename(systemFilename)
	if errLoadSystem != nil {
		panic("Datei fehlt")
	}

	allItems, err2 := pf2e.NewAllPacks(sys.Packs, systemPath)
	if err2 != nil {
		log.Fatal(err2)
		return
	}

	fmt.Println(len(allItems))
}
