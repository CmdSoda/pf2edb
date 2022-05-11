package pf2e

type Item struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Data struct {
		Description struct {
			Value string `json:"value"`
		} `json:"description"`
	} `json:"data"`
}
