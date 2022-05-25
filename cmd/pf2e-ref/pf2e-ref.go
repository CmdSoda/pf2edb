package main

import (
	"fmt"
	"github.com/CmdSoda/pf2edb/internal/pf2e"
	"github.com/alexflint/go-arg"
	"log"
	"os"
	"path/filepath"
)

// usage: pf2e-ref.exe [-c type] <search-string>
func main() {
	fmt.Println("pf2e-ref.exe V1.0.1")
	var args struct {
		Type            string `arg:"-t"`
		ShowDescription bool   `arg:"-d"`
		Exact           bool   `arg:"-e"`
		Interactive     bool   `arg:"-i"`
		Search          string `arg:"positional, required"`
	}
	arg.MustParse(&args)

	config := "./pf2e-ref.config.json"
	configPath := ""

	if _, err := os.Stat(config); err == nil {
		configPath = config
	} else {
		path, errExe := os.Executable()
		if errExe != nil {
			log.Println(errExe)
		}
		configPath = filepath.Dir(path) + "./pf2e-ref.config.json"
	}

	//configPath := "./pf2e-ref.config.json"
	fmt.Println("loading config " + configPath)

	ctx := pf2e.NewContext()
	err := ctx.Config.Load(configPath)
	if err != nil {
		panic("verdammt")
	}

	if _, err := os.Stat(ctx.Config.DataFolder); os.IsNotExist(err) {
		log.Fatal(ctx.Config.DataFolder + " does not exist")
		return
	}

	systemsFolder := ctx.Config.DataFolder + "\\systems\\"

	if _, err := os.Stat(systemsFolder); os.IsNotExist(err) {
		log.Fatal(systemsFolder + " does not exist")
		return
	}

	systemNamePath := systemsFolder + ctx.Config.SystemName

	if _, err := os.Stat(systemNamePath); os.IsNotExist(err) {
		log.Fatal(systemNamePath + " does not exist")
		return
	}
	systemFilename := systemNamePath + "\\system.json"
	sys, errLoadSystem := pf2e.NewSystemFromFilename(systemFilename)
	if errLoadSystem != nil {
		panic("Datei fehlt")
	}

	allItems, err2 := pf2e.NewAllPacks(sys.Packs, systemNamePath)
	if err2 != nil {
		log.Fatal(err2)
		return
	}

	translationModuleCompendiumFolder := ctx.Config.DataFolder + "\\modules\\" + ctx.Config.TranslationModule + "\\compendium"
	tl, err3 := pf2e.NewAllTranslations(translationModuleCompendiumFolder)
	if err3 != nil {
		log.Fatal(err3)
		return
	}

	pf2e.DoTranslate(&allItems, &tl)
	pf2e.Search(allItems, args.Search, args.ShowDescription, args.Exact, args.Interactive)
}
