package pf2e

type TranslationEntryData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ConverterName struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

type TranslationConverterEntryData struct {
	Name        ConverterName
	Description string `json:"description"`
}

type TranslationData struct {
	Label string `json:"label"`
	//Mapping TranslationConverterEntryData   `json:"mapping"`
	Entries map[string]TranslationEntryData `json:"entries"`
}

type AllTranslations []TranslationData
