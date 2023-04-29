package sonolus

type LevelItem struct {
	Name          string                  `json:"name"`
	Version       int                     `json:"version"`
	Rating        int                     `json:"rating"`
	Title         string                  `json:"title"`
	Artists       string                  `json:"artists"`
	Author        string                  `json:"author"`
	Engine        EngineItem              `json:"engine"`
	UseSkin       UseItem[SkinItem]       `json:"useSkin"`
	UseBackground UseItem[BackgroundItem] `json:"useBackground"`
	UseEffect     UseItem[EffectItem]     `json:"useEffect"`
	UseParticle   UseItem[ParticleItem]   `json:"useParticle"`
	Cover         SRL                     `json:"cover"`
	Bgm           SRL                     `json:"bgm"`
	Preview       *SRL                    `json:"preview,omitempty"`
	Data          SRL                     `json:"data"`
}

type UseItem[T any] struct {
	UseDefault bool `json:"useDefault"`
	Item       T    `json:"item,omitempty"`
}
type LevelDataEntity struct {
	Archetype int                  `json:"archetype"`
	Data      *LevelDataEntityData `json:"data,omitempty"`
}

type LevelDataEntityData struct {
	Index  int   `json:"index"`
	Values []float64 `json:"values"`
}

type LevelData struct {
	BgmOffset float64           `json:"bgmOffset"`
	Entities  []LevelDataEntity `json:"entities"`
}
