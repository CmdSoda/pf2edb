package pf2e

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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
			ret = append(ret, &items)
		}
	}
	return ret, nil
}

func (p Pack) GetItems(systemRoot string) (PackItems, error) {
	var itemList []*Item
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
		item.modulePath = p.System + "." + p.Name
		itemList = append(itemList, &item)
	}
	if errScan := scanner.Err(); errScan != nil {
		return itemList, errScan
	}
	return itemList, nil
}

func NewAllTranslations(compendiumFolder string) (AllTranslations, error) {
	tl := AllTranslations{}

	files, err := ioutil.ReadDir(compendiumFolder)
	if err != nil {
		return tl, err
	}

	for _, file := range files {
		if !file.IsDir() {
			compendiumFilepath := compendiumFolder + "\\" + file.Name()
			ext := filepath.Ext(compendiumFilepath)
			if ext == ".json" {
				content, err := ioutil.ReadFile(compendiumFilepath)
				if err != nil {
					return tl, err
				}
				td := TranslationData{}
				err = json.Unmarshal(content, &td)
				if err != nil {
					return tl, err
				}
				tl = append(tl, td)
			}
		}
	}

	return tl, nil
}

func DoTranslate(allPacks *AllPacks, allTranslations *AllTranslations) {
	for translationIndex := 0; translationIndex < len(*allTranslations); translationIndex++ {
		td := (*allTranslations)[translationIndex]
		for key, _ := range td.Entries {
			ted := td.Entries[key]
			for packItemsListIndex := 0; packItemsListIndex < len(*allPacks); packItemsListIndex++ {
				pi := (*allPacks)[packItemsListIndex]
				for itemsListIndex := 0; itemsListIndex < len(*pi); itemsListIndex++ {
					i := (*pi)[itemsListIndex]
					if i.Name == key && i.Name != ted.Name {
						i.NameTranslation = ted.Name
						i.Data.Description.Value = ted.Description
					}
				}
			}
		}
	}
}

func Search(allPacks AllPacks, search string, showDesc bool) {
	for packItemsListIndex := 0; packItemsListIndex < len(allPacks); packItemsListIndex++ {
		pi := allPacks[packItemsListIndex]
		for itemsListIndex := 0; itemsListIndex < len(*pi); itemsListIndex++ {
			i := (*pi)[itemsListIndex]
			searchLower := strings.ToLower(search)
			if strings.ToLower(i.Name) == searchLower ||
				strings.ToLower(i.NameTranslation) == searchLower ||
				strings.Contains(strings.ToLower(i.Name), searchLower) ||
				strings.Contains(strings.ToLower(i.NameTranslation), searchLower) {
				if i.GetCleanTranslation() != "" {
					fmt.Println("@Compendium[" + i.modulePath + "." + i.ID + "]{" + i.GetCleanTranslation() + "}")
				} else {
					fmt.Println("@Compendium[" + i.modulePath + "." + i.ID + "]{" + i.Name + "}")
				}
				if showDesc {
					fmt.Println(i.Data.Description.Value)
					fmt.Println()
				}
			}
		}
	}
}
