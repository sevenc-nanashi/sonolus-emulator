package processor

type MemoryPermission struct {
	Read  bool
	Write bool
}

type EntityMemories struct {
	Memory []float64
	Input  []float64
}

func (p *Processor) initMemories() {
	p.memories = make(map[int][]float64)
	p.memories[0] = make([]float64, 4096)                                    // Level Memory
	p.memories[1] = make([]float64, 4096)                                    // Level Data
	p.memories[2] = make([]float64, len(p.data.EngineConfiguration.Options)) // []Level Option
	p.memories[3] = make([]float64, 16)                                      // Level Transform
	p.memories[4] = make([]float64, 16)                                      // Level Background
	p.memories[5] = make([]float64, 80)                                      // Level UI
	p.memories[6] = make([]float64, 6*len(p.data.EngineData.Buckets))        // []Level Bucket
	p.memories[7] = make([]float64, 12)                                      // Level Score
	p.memories[8] = make([]float64, 6)                                       // Level Life
	p.memories[9] = make([]float64, 10)                                      // Level UI Configuration
	p.memories[10] = make([]float64, 32*len(p.data.LevelData.Entities))      // []Entity Info
	p.memories[11] = make([]float64, 32*len(p.data.LevelData.Entities))      // []Entity Data
	p.memories[12] = make([]float64, 32*len(p.data.LevelData.Entities))      // []Entity Shared Memory
	p.memories[20] = make([]float64, 3)                                      // Entity Info
	p.memories[21] = make([]float64, 32)                                     // Entity Memory
	p.memories[22] = make([]float64, 32)                                     // Entity Data
	p.memories[23] = make([]float64, 4)                                      // Entity Input
	p.memories[24] = make([]float64, 32)                                     // Entity Shared Memory
	p.memories[30] = make([]float64, 3*len(p.data.EngineData.Archetypes))    // []Archetype Life

	rom := make([]float64, len(p.data.EngineRom))
	for i, value := range p.data.EngineRom {
		rom[i] = float64(value)
	}
	p.memories[50] = rom // Engine Rom

	p.memories[100] = make([]float64, 4096) // Temporary Memory
	p.memories[101] = make([]float64, 15)   // Temporary Data

	p.memories[1][2] = p.config.AspectRatio
	p.EntityMemories = map[int]EntityMemories{}
  for entityIndex := range p.data.LevelData.Entities {
    p.EntityMemories[entityIndex] = EntityMemories{
      Memory: make([]float64, 32),
      Input:  make([]float64, 4),
    }
  }
}

func (p *Processor) prepareMemory(entityIndex int, status Status) {
	p.memories[100] = make([]float64, 4096) // Temporary Memory
	p.memories[21] = p.EntityMemories[entityIndex].Memory
	p.memories[23] = p.EntityMemories[entityIndex].Input
	switch status {
	case StatusPreprocess:
		p.memoryPermissions = map[int]MemoryPermission{
			0:   {Read: true, Write: true},
			1:   {Read: true, Write: true},
			2:   {Read: true, Write: true},
			3:   {Read: true, Write: true},
			4:   {Read: true, Write: true},
			5:   {Read: true, Write: true},
			6:   {Read: true, Write: true},
			7:   {Read: true, Write: true},
			8:   {Read: true, Write: true},
			9:   {Read: true, Write: true},
			10:  {Read: true, Write: false},
			11:  {Read: true, Write: true},
			12:  {Read: true, Write: true},
			20:  {Read: true, Write: false},
			21:  {Read: true, Write: true},
			22:  {Read: true, Write: true},
			23:  {Read: true, Write: true},
			24:  {Read: true, Write: true},
			30:  {Read: true, Write: true},
			50:  {Read: true, Write: false},
			100: {Read: true, Write: true},
			101: {Read: false, Write: false},
		}
	case StatusSpawnOrder:
		p.memoryPermissions = map[int]MemoryPermission{
			0:   {Read: true, Write: false},
			1:   {Read: true, Write: false},
			2:   {Read: true, Write: false},
			3:   {Read: true, Write: false},
			4:   {Read: true, Write: false},
			5:   {Read: true, Write: false},
			6:   {Read: true, Write: false},
			7:   {Read: true, Write: false},
			8:   {Read: true, Write: false},
			9:   {Read: true, Write: false},
			10:  {Read: true, Write: false},
			11:  {Read: true, Write: false},
			12:  {Read: true, Write: false},
			20:  {Read: true, Write: false},
			21:  {Read: true, Write: true},
			22:  {Read: true, Write: false},
			23:  {Read: true, Write: true},
			24:  {Read: true, Write: false},
			30:  {Read: true, Write: false},
			50:  {Read: true, Write: false},
			100: {Read: true, Write: true},
			101: {Read: false, Write: false},
		}
	case StatusShouldSpawn:
		p.memoryPermissions = map[int]MemoryPermission{
			0:   {Read: true, Write: false},
			1:   {Read: true, Write: false},
			2:   {Read: true, Write: false},
			3:   {Read: true, Write: false},
			4:   {Read: true, Write: false},
			5:   {Read: true, Write: false},
			6:   {Read: true, Write: false},
			7:   {Read: true, Write: false},
			8:   {Read: true, Write: false},
			9:   {Read: true, Write: false},
			10:  {Read: true, Write: false},
			11:  {Read: true, Write: false},
			12:  {Read: true, Write: false},
			20:  {Read: true, Write: false},
			21:  {Read: true, Write: true},
			22:  {Read: true, Write: false},
			23:  {Read: true, Write: true},
			24:  {Read: true, Write: false},
			30:  {Read: true, Write: false},
			50:  {Read: true, Write: false},
			100: {Read: true, Write: true},
			101: {Read: false, Write: false},
		}
	case StatusInitialize:
		p.memoryPermissions = map[int]MemoryPermission{
			0:   {Read: true, Write: false},
			1:   {Read: true, Write: false},
			2:   {Read: true, Write: false},
			3:   {Read: true, Write: false},
			4:   {Read: true, Write: false},
			5:   {Read: true, Write: false},
			6:   {Read: true, Write: false},
			7:   {Read: true, Write: false},
			8:   {Read: true, Write: false},
			9:   {Read: true, Write: false},
			10:  {Read: true, Write: false},
			11:  {Read: true, Write: false},
			12:  {Read: true, Write: false},
			20:  {Read: true, Write: false},
			21:  {Read: true, Write: true},
			22:  {Read: true, Write: false},
			23:  {Read: true, Write: true},
			24:  {Read: true, Write: false},
			30:  {Read: true, Write: false},
			50:  {Read: true, Write: false},
			100: {Read: true, Write: true},
			101: {Read: false, Write: false},
		}
	case StatusUpdateSequential:
		p.memoryPermissions = map[int]MemoryPermission{
			0:   {Read: true, Write: true},
			1:   {Read: true, Write: false},
			2:   {Read: true, Write: false},
			3:   {Read: true, Write: true},
			4:   {Read: true, Write: true},
			5:   {Read: true, Write: false},
			6:   {Read: true, Write: false},
			7:   {Read: true, Write: false},
			8:   {Read: true, Write: false},
			9:   {Read: true, Write: false},
			10:  {Read: true, Write: false},
			11:  {Read: true, Write: false},
			12:  {Read: true, Write: true},
			20:  {Read: true, Write: false},
			21:  {Read: true, Write: true},
			22:  {Read: true, Write: false},
			23:  {Read: true, Write: true},
			24:  {Read: true, Write: true},
			30:  {Read: true, Write: false},
			50:  {Read: true, Write: false},
			100: {Read: true, Write: true},
			101: {Read: false, Write: false},
		}
	case StatusTouch:
		p.memoryPermissions = map[int]MemoryPermission{
			0:   {Read: true, Write: true},
			1:   {Read: true, Write: false},
			2:   {Read: true, Write: false},
			3:   {Read: true, Write: true},
			4:   {Read: true, Write: true},
			5:   {Read: true, Write: false},
			6:   {Read: true, Write: false},
			7:   {Read: true, Write: false},
			8:   {Read: true, Write: false},
			9:   {Read: true, Write: false},
			10:  {Read: true, Write: false},
			11:  {Read: true, Write: false},
			12:  {Read: true, Write: true},
			20:  {Read: true, Write: false},
			21:  {Read: true, Write: true},
			22:  {Read: true, Write: false},
			23:  {Read: true, Write: true},
			24:  {Read: true, Write: true},
			30:  {Read: true, Write: false},
			50:  {Read: true, Write: false},
			100: {Read: true, Write: true},
			101: {Read: true, Write: false},
		}
	case StatusUpdateParallel:
		p.memoryPermissions = map[int]MemoryPermission{
			0:   {Read: true, Write: false},
			1:   {Read: true, Write: false},
			2:   {Read: true, Write: false},
			3:   {Read: true, Write: false},
			4:   {Read: true, Write: false},
			5:   {Read: true, Write: false},
			6:   {Read: true, Write: false},
			7:   {Read: true, Write: false},
			8:   {Read: true, Write: false},
			9:   {Read: true, Write: false},
			10:  {Read: true, Write: false},
			11:  {Read: true, Write: false},
			12:  {Read: true, Write: false},
			20:  {Read: true, Write: false},
			21:  {Read: true, Write: true},
			22:  {Read: true, Write: false},
			23:  {Read: true, Write: true},
			24:  {Read: true, Write: false},
			30:  {Read: true, Write: false},
			50:  {Read: true, Write: false},
			100: {Read: true, Write: true},
			101: {Read: false, Write: false},
		}
	case StatusTerminate:
		p.memoryPermissions = map[int]MemoryPermission{
			0:   {Read: true, Write: false},
			1:   {Read: true, Write: false},
			2:   {Read: true, Write: false},
			3:   {Read: true, Write: false},
			4:   {Read: true, Write: false},
			5:   {Read: true, Write: false},
			6:   {Read: true, Write: false},
			7:   {Read: true, Write: false},
			8:   {Read: true, Write: false},
			9:   {Read: true, Write: false},
			10:  {Read: true, Write: false},
			11:  {Read: true, Write: false},
			12:  {Read: true, Write: false},
			20:  {Read: true, Write: false},
			21:  {Read: true, Write: true},
			22:  {Read: true, Write: false},
			23:  {Read: true, Write: true},
			24:  {Read: true, Write: false},
			30:  {Read: true, Write: false},
			50:  {Read: true, Write: false},
			100: {Read: true, Write: true},
			101: {Read: false, Write: false},
		}
	default:
		panic("unknown status")
	}
}
