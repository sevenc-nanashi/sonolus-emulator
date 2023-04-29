package sonolus

type SRL struct {
	Type ResourceType `json:"type"`
	Hash string       `json:"hash"`
	Url  string       `json:"url"`
}

type ResourceType string
