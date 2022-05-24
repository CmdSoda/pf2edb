package main

import (
	"fmt"
	"github.com/CmdSoda/pf2edb/internal/pf2e"
	"log"
)

func main() {
	fmt.Println("pf2e-ref.exe V1.0.1")

	/*
		path, errExe := os.Executable()
		if errExe != nil {
			log.Println(errExe)
		}
		configPath := filepath.Dir(path) + "./pf2e-ref.config.json"

	*/
	configPath := "./pf2e-ref.config.json"

	fmt.Println("loading config " + configPath)

	ctx := pf2e.NewContext()
	err := ctx.Config.Load(configPath)
	if err != nil {
		panic("verdammt")
	}

	translationModuleCompendiumFolder := ctx.Config.DataFolder + "\\modules\\" + ctx.Config.TranslationModule + "\\compendium"
	err3 := pf2e.RemoveAllEnglishWords(translationModuleCompendiumFolder)
	if err3 != nil {
		log.Fatal(err3)
		return
	}
}
