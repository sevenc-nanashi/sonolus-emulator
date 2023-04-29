package processor

import (
	"math/rand"
	"time"
)

type Processor struct {
	data ProcessorData

	Entities   []Entity
	spawnQueue []SpawnQueueItem
	DrawCalls  []DrawCall
	DebugLog   float64

	memories          map[int][]float64
	memoryPermissions map[int]MemoryPermission
	EntityMemories    map[int]EntityMemories

	config ProcessorConfig
}

type ProcessorConfig struct {
	AspectRatio float64
}

func Init(config ProcessorConfig) Processor {
	processor := Processor{
		config: config,
	}
	rand.Seed(time.Now().UnixNano())

	return processor
}

func (p *Processor) Load(serverUrl string, levelName string) error {
	var processorData ProcessorData = ProcessorData{}

	err := processorData.loadLevel(serverUrl, levelName)
	if err != nil {
		return err
	}

	err = processorData.loadEngine(serverUrl, processorData.LevelItem.Engine)
	if err != nil {
		return err
	}

	p.data = processorData

	p.initMemories()
	p.loadEntities()

	return nil
}
