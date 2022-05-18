package pf2e

import "strings"

type Item struct {
	ID              string `json:"_id"`
	Name            string `json:"name"`
	NameTranslation string `json:"nametranslation"`
	Type            string `json:"type"`
	modulePath      string
	Data            struct {
		Description struct {
			Value string `json:"value"`
		} `json:"description"`
	} `json:"data"`
}

func (i Item) GetCleanTranslation() string {
	cleanIndex := strings.Index(i.NameTranslation, "/")
	if cleanIndex != -1 {
		return i.NameTranslation[0:cleanIndex]
	} else {
		return i.NameTranslation
	}
}
