package processor

import (
	"github.com/sevenc-nanashi/sonolus-emulator/pkg/sonolus"

	log "github.com/sirupsen/logrus"
)

type Entity struct {
	Archetype  Archetype
	Values     []float64
	SpawnOrder float64

	initialized bool
}

type SpawnQueueItem struct {
	Scripts sonolus.EngineDataScript
	Data    []float64
}

type Archetype struct {
	Processor *Processor
	Index     int
	Scripts   sonolus.EngineDataScript
}

func (p *Processor) loadEntities() Processor {
	var entities = make([]Entity, len(p.data.LevelData.Entities))
	for index, entity := range p.data.LevelData.Entities {
		archetype := p.data.EngineData.Archetypes[entity.Archetype]
		values := make([]float64, 32)
		if archetype.Data != nil {
			for i, value := range archetype.Data.Values {
				values[archetype.Data.Index+i] = value
			}
		}
		if entity.Data != nil {
			for i, value := range entity.Data.Values {
				values[entity.Data.Index+i] = value
			}
		}
		entities[index] = Entity{
			Archetype: Archetype{
				Processor: p,
				Index:     entity.Archetype,
				Scripts:   p.data.EngineData.Scripts[archetype.Script],
			},
			Values:      values,
			SpawnOrder:  0,
			initialized: false,
		}
	}
	p.Entities = entities
	return *p
}

func (a *Archetype) Preprocess(entityIndex int) {
	log.Infof("Preprocessing entity %d", entityIndex)
	if a.Scripts.Preprocess == nil {
		return
	}
	a.Processor.Execute(StatusPreprocess, entityIndex, a.Scripts.Preprocess.Index)
}

func (a *Archetype) SpawnOrder(entityIndex int) float64 {
	// log.Infof("Calculating spawn order for entity %d", entityIndex)
	if a.Scripts.SpawnOrder == nil {
		return 0
	}
	return a.Processor.Execute(StatusSpawnOrder, entityIndex, a.Scripts.SpawnOrder.Index)
}

func (a *Archetype) Initialize(entityIndex int) {
	// log.Infof("Initializing entity %d", entityIndex)
	if a.Scripts.Initialize == nil {
		return
	}
	a.Processor.Execute(StatusInitialize, entityIndex, a.Scripts.Initialize.Index)
}

func (a *Archetype) UpdateSequential(entityIndex int) {
	// log.Infof("Running sequential update for entity %d", entityIndex)
	if a.Scripts.UpdateSequential == nil {
		return
	}
	a.Processor.Execute(StatusUpdateSequential, entityIndex, a.Scripts.UpdateSequential.Index)
}

func (a *Archetype) Touch(entityIndex int) {
	// log.Infof("Running touch for entity %d", entityIndex)
	if a.Scripts.Touch == nil {
		return
	}
	a.Processor.Execute(StatusTouch, entityIndex, a.Scripts.Touch.Index)
}
