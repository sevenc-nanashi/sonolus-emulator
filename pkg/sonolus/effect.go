package sonolus

type EffectItem struct {
	Name      string `json:"name"`
	Version   int    `json:"version"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Author    string `json:"author"`
	Thumbnail SRL    `json:"thumbnail"`
	Data      SRL    `json:"data"`
	Audio     SRL    `json:"audio"`
}
