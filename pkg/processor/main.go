package processor

type Processor struct {
	Data ProcessorData

	Entities []Entity

	Memories map[int][]float64
  MemoryPermissions map[int]MemoryPermission
}

func Load(serverUrl string, levelName string) (Processor, error) {
	var processorData ProcessorData = ProcessorData{}

	err := processorData.loadLevel(serverUrl, levelName)
	if err != nil {
		return Processor{}, err
	}

	err = processorData.loadEngine(serverUrl, processorData.LevelItem.Engine)
	if err != nil {
		return Processor{}, err
	}

	processor := Processor{
		Data: processorData,
	}

	processor.initMemories()
	processor.loadEntities()

	return processor, nil
}
