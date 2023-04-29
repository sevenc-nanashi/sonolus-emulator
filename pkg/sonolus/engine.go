package sonolus

import (
	"github.com/byrnedo/pjson"
)

type EngineItem struct {
	Name          string         `json:"name"`
	Version       int            `json:"version"`
	Title         string         `json:"title"`
	Subtitle      string         `json:"subtitle"`
	Author        string         `json:"author"`
	Skin          SkinItem       `json:"skin"`
	Background    BackgroundItem `json:"background"`
	Effect        EffectItem     `json:"effect"`
	Particle      ParticleItem   `json:"particle"`
	Thumbnail     SRL            `json:"thumbnail"`
	Data          SRL            `json:"data"`
	Rom           *SRL           `json:"rom,omitempty"`
	Configuration SRL            `json:"configuration"`
}

type EngineDataNode struct {
	Value float64 `json:"value,omitempty"`
	Func  string  `json:"func,omitempty"`
	Args  []int   `json:"args,omitempty"`
}

type EngineDataScriptCallback struct {
	Index int  `json:"index"`
	Order *int `json:"order,omitempty"`
}

type EngineDataScript struct {
	Preprocess       *EngineDataScriptCallback `json:"preprocess,omitempty"`
	SpawnOrder       *EngineDataScriptCallback `json:"spawnOrder,omitempty"`
	ShouldSpawn      *EngineDataScriptCallback `json:"shouldSpawn,omitempty"`
	Initialize       *EngineDataScriptCallback `json:"initialize,omitempty"`
	UpdateSequential *EngineDataScriptCallback `json:"updateSequential,omitempty"`
	Touch            *EngineDataScriptCallback `json:"touch,omitempty"`
	UpdateParallel   *EngineDataScriptCallback `json:"updateParallel,omitempty"`
	Terminate        *EngineDataScriptCallback `json:"terminate,omitempty"`
}

type EngineData struct {
	Buckets    []EngineDataBucket    `json:"buckets"`
	Archetypes []EngineDataArchetype `json:"archetypes"`
	Scripts    []EngineDataScript    `json:"scripts"`
	Nodes      []EngineDataNode      `json:"nodes"`
}

type EngineDataBucket struct {
	Sprites []EngineDataSprite `json:"sprites"`
}

type EngineDataSprite struct {
	Id         int  `json:"id"`
	FallbackId *int `json:"fallbackId,omitempty"`
	X          int  `json:"x"`
	Y          int  `json:"y"`
	W          int  `json:"w"`
	H          int  `json:"h"`
	Rotation   int  `json:"rotation"`
}

type EngineDataArchetype struct {
	Script int `json:"script"`
	Data   *struct {
		Index  int       `json:"index"`
		Values []float64 `json:"values"`
	} `json:"data,omitempty"`
	Input bool `json:"input,omitempty"`
}

type EngineConfigurationUI struct {
	Scope                     *string                       `json:"scope,omitempty"`
	PrimaryMetric             string                        `json:"primaryMetric"`
	SecondaryMetric           string                        `json:"secondaryMetric"`
	MenuVisibility            EngineConfigurationVisibility `json:"menuVisibility"`
	JudgmentVisibility        EngineConfigurationVisibility `json:"judgmentVisibility"`
	ComboVisibility           EngineConfigurationVisibility `json:"comboVisibility"`
	PrimaryMetricVisibility   EngineConfigurationVisibility `json:"primaryMetricVisibility"`
	SecondaryMetricVisibility EngineConfigurationVisibility `json:"secondaryMetricVisibility"`
	JudgmentAnimation         EngineConfigurationAnimation  `json:"judgmentAnimation"`
	ComboAnimation            EngineConfigurationAnimation  `json:"comboAnimation"`
	JudgmentErrorStyle        string                        `json:"judgmentErrorStyle"`
	JudgmentErrorPlacement    string                        `json:"judgmentErrorPlacement"`
	JudgmentErrorMin          int                           `json:"judgmentErrorMin"`
}

type EngineConfigurationMetric string

const (
	EngineConfigurationMetricArcade       EngineConfigurationMetric = "arcade"
	EngineConfigurationMetricAccuracy     EngineConfigurationMetric = "accuracy"
	EngineConfigurationMetricLife         EngineConfigurationMetric = "life"
	EngineConfigurationMetricPerfectRate  EngineConfigurationMetric = "perfectRate"
	EngineConfigurationMetricErrorHeatmap EngineConfigurationMetric = "errorHeatmap"
)

type EngineConfigurationVisibility struct {
	Scale float64 `json:"scale"`
	Alpha float64 `json:"alpha"`
}

type EngineConfigurationAnimation struct {
	Scale EngineConfigurationAnimationTween `json:"scale"`
	Alpha EngineConfigurationAnimationTween `json:"alpha"`
}

type EngineConfigurationAnimationTween struct {
	From     float64 `json:"from"`
	To       float64 `json:"to"`
	Duration int     `json:"duration"`
	Ease     string  `json:"ease"`
}

type EngineConfigurationJudgmentErrorStyle string

const (
	EngineConfigurationJudgmentErrorStyleNone          EngineConfigurationJudgmentErrorStyle = "none"
	EngineConfigurationJudgmentErrorStylePlus          EngineConfigurationJudgmentErrorStyle = "plus"
	EngineConfigurationJudgmentErrorStyleMinus         EngineConfigurationJudgmentErrorStyle = "minus"
	EngineConfigurationJudgmentErrorStyleArrowUp       EngineConfigurationJudgmentErrorStyle = "arrowUp"
	EngineConfigurationJudgmentErrorStyleArrowDown     EngineConfigurationJudgmentErrorStyle = "arrowDown"
	EngineConfigurationJudgmentErrorStyleArrowLeft     EngineConfigurationJudgmentErrorStyle = "arrowLeft"
	EngineConfigurationJudgmentErrorStyleArrowRight    EngineConfigurationJudgmentErrorStyle = "arrowRight"
	EngineConfigurationJudgmentErrorStyleTriangleUp    EngineConfigurationJudgmentErrorStyle = "triangleUp"
	EngineConfigurationJudgmentErrorStyleTriangleDown  EngineConfigurationJudgmentErrorStyle = "triangleDown"
	EngineConfigurationJudgmentErrorStyleTriangleLeft  EngineConfigurationJudgmentErrorStyle = "triangleLeft"
	EngineConfigurationJudgmentErrorStyleTriangleRight EngineConfigurationJudgmentErrorStyle = "triangleRight"
)

type EngineConfigurationJudgmentErrorPlacement string

const (
	EngineConfigurationJudgmentErrorPlacementBoth  EngineConfigurationJudgmentErrorPlacement = "both"
	EngineConfigurationJudgmentErrorPlacementLeft  EngineConfigurationJudgmentErrorPlacement = "left"
	EngineConfigurationJudgmentErrorPlacementRight EngineConfigurationJudgmentErrorPlacement = "right"
)

type EngineConfiguration struct {
	Options []EngineConfigurationOption `json:"options"`
	UI      EngineConfigurationUI       `json:"ui"`
}

type EngineConfigurationOption struct{}

func (u EngineConfigurationOption) Field() string {
	return "type"
}

func (u EngineConfigurationOption) Variants() []pjson.Variant {
	return []pjson.Variant{
		EngineConfigurationSliderOption{},
		EngineConfigurationToggleOption{},
		EngineConfigurationSelectOption{},
	}
}

type EngineConfigurationSliderOption struct {
	Name     string  `json:"name"`
	Standard *bool   `json:"standard,omitempty"`
	Scope    *string `json:"scope,omitempty"`
	Type     string  `json:"type"`
	Def      float64 `json:"def"`
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Step     float64 `json:"step"`
	Unit     *string `json:"unit,omitempty"`
}

func (u EngineConfigurationSliderOption) Variant() string {
	return "slider"
}

type EngineConfigurationToggleOption struct {
	Name     string  `json:"name"`
	Standard *bool   `json:"standard,omitempty"`
	Scope    *string `json:"scope,omitempty"`
	Type     string  `json:"type"`
	Def      int     `json:"def"`
}

func (u EngineConfigurationToggleOption) Variant() string {
	return "toggle"
}

type EngineConfigurationSelectOption struct {
	Name     string   `json:"name"`
	Standard *bool    `json:"standard,omitempty"`
	Scope    *string  `json:"scope,omitempty"`
	Type     string   `json:"type"`
	Def      int      `json:"def"`
	Values   []string `json:"values"`
}

func (u EngineConfigurationSelectOption) Variant() string {
	return "select"
}
