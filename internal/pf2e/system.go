package pf2e

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

type PackItems []*Item
type AllPacks []*PackItems

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
