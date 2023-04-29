package processor

import "github.com/sevenc-nanashi/sonolus-emulator/pkg/sonolus"

type Entity struct {
	Archetype Archetype
	Values    []float64
  SpawnOrder int
}

type Archetype struct {
	Processor *Processor
	Index     int
	Scripts   sonolus.EngineDataScript
}

func (p *Processor) loadEntities() Processor {
	var entities = make([]Entity, len(p.Data.LevelData.Entities))
	for index, entity := range p.Data.LevelData.Entities {
		archetype := p.Data.EngineData.Archetypes[entity.Archetype]
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
				Scripts:   p.Data.EngineData.Scripts[archetype.Script],
			},
			Values: values,
      SpawnOrder: 0,
		}
	}
	p.Entities = entities
	return *p
}

func (a *Archetype) Preprocess(entityIndex int) {
	if a.Scripts.Preprocess == nil {
		return
	}
	a.Processor.Execute(StatusPreprocess, entityIndex, a.Scripts.Preprocess.Index)
}
