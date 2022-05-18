package pf2e

type TranslationEntryData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TranslationData struct {
	Label   string                          `json:"label"`
	Entries map[string]TranslationEntryData `json:"entries"`
}

type AllTranslations []TranslationData