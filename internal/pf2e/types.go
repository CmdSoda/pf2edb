package pf2e

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Pack struct {
	Name    string `json:"name"`
	Label   string `json:"label"`
	System  string `json:"system"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	Private bool   `json:"private,omitempty"`
	Folder  string `json:"folder,omitempty"`
	Module  string `json:"module,omitempty"`
}

type InvalidDbCategoryError struct{}

func (ic InvalidDbCategoryError) Error() string {
	return "Invalid db category"
}

type PackItems []Item
type AllPacks []PackItems

func NewAllPacks(packs []Pack, systemPath string) (AllPacks, error) {
	var ret AllPacks
	for i := 0; i < len(packs); i++ {
		p := packs[i]
		if p.Type == "Item" {
			items, err := p.GetItems(systemPath)
			if err != nil {
				panic(err)
				return ret, err
			}
			ret = append(ret, items)
		}
	}
	return ret, nil
}

func (p Pack) GetItems(systemRoot string) (PackItems, error) {
	var itemList []Item
	if p.Type != "Item" {
		return itemList, InvalidDbCategoryError{}
	}
	filepath := systemRoot + "\\" + p.Path
	file, err := os.Open(filepath)
	if err != nil {
		return itemList, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		item := Item{}
		err = json.Unmarshal([]byte(line), &item)
		if err != nil {
			return itemList, err
		}
		itemList = append(itemList, item)
	}
	if errScan := scanner.Err(); errScan != nil {
		return itemList, errScan
	}
	return itemList, nil
}

type Language struct {
	Lang string `json:"lang"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type System struct {
	Name                  string     `json:"name"`
	Title                 string     `json:"title"`
	Description           string     `json:"description"`
	License               string     `json:"license"`
	Version               string     `json:"version"`
	MinimumCoreVersion    string     `json:"minimumCoreVersion"`
	CompatibleCoreVersion string     `json:"compatibleCoreVersion"`
	Author                string     `json:"author"`
	Esmodules             []string   `json:"esmodules"`
	Scripts               []string   `json:"scripts"`
	Styles                []string   `json:"styles"`
	Packs                 []Pack     `json:"packs"`
	Languages             []Language `json:"languages"`
	Socket                bool       `json:"socket"`
	TemplateVersion       int        `json:"templateVersion"`
	Initiative            string     `json:"initiative"`
	GridDistance          int        `json:"gridDistance"`
	GridUnits             string     `json:"gridUnits"`
	PrimaryTokenAttribute string     `json:"primaryTokenAttribute"`
	URL                   string     `json:"url"`
	Bugs                  string     `json:"bugs"`
	Changelog             string     `json:"changelog"`
	Manifest              string     `json:"manifest"`
	Download              string     `json:"download"`
}

func NewSystemFromFilename(filename string) (System, error) {
	content, err := ioutil.ReadFile(filename)
	sys := System{}
	if err != nil {
		return sys, err
	}
	err = json.Unmarshal(content, &sys)
	if err != nil {
		return sys, err
	}
	return sys, nil
}
