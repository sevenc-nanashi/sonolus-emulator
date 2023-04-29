package sonolus

type BackgroundItem struct {
	Name          string `json:"name"`
	Version       int    `json:"version"`
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	Author        string `json:"author"`
	Thumbnail     SRL    `json:"thumbnail"`
	Data          SRL    `json:"data"`
	Image         SRL    `json:"image"`
	Configuration SRL    `json:"configuration"`
}
