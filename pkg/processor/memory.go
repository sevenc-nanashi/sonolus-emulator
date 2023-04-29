package processor

type MemoryPermission struct {
	Read  bool
	Write bool
}

func (p *Processor) initMemories() {
	p.Memories = make(map[int][]float64)
	p.Memories[0] = make([]float64, 4096)                                    // Level Memory
	p.Memories[1] = make([]float64, 4096)                                    // Level Data
	p.Memories[2] = make([]float64, len(p.Data.EngineConfiguration.Options)) // []Level Option
	p.Memories[3] = make([]float64, 16)                                      // Level Transform
	p.Memories[4] = make([]float64, 16)                                      // Level Background
	p.Memories[5] = make([]float64, 80)                                      // Level UI
	p.Memories[6] = make([]float64, 6*len(p.Data.EngineData.Buckets))        // []Level Bucket
	p.Memories[7] = make([]float64, 12)                                      // Level Score
	p.Memories[8] = make([]float64, 6)                                       // Level Life
	p.Memories[9] = make([]float64, 10)                                      // Level UI Configuration
	p.Memories[10] = make([]float64, 32*len(p.Data.LevelData.Entities))      // []Entity Info
	p.Memories[11] = make([]float64, 32*len(p.Data.LevelData.Entities))      // []Entity Data
	p.Memories[12] = make([]float64, 32*len(p.Data.LevelData.Entities))      // []Entity Shared Memory
	p.Memories[20] = make([]float64, 3)                                      // Entity Info
	p.Memories[21] = make([]float64, 32)                                     // Entity Memory
	p.Memories[22] = make([]float64, 32)                                     // Entity Data
	p.Memories[23] = make([]float64, 4)                                      // Entity Input
	p.Memories[24] = make([]float64, 32)                                     // Entity Shared Memory
	p.Memories[30] = make([]float64, 3*len(p.Data.EngineData.Archetypes))    // []Archetype Life

	rom := make([]float64, len(p.Data.EngineRom))
	for i, value := range p.Data.EngineRom {
		rom[i] = float64(value)
	}
	p.Memories[50] = rom // Engine Rom

	p.Memories[100] = make([]float64, 4096) // Temporary Memory
	p.Memories[101] = make([]float64, 15)   // Temporary Data
}

func (p *Processor) prepareMemory(entityIndex int, status Status) {
  p.Memories[20] = p.Memories[10][entityIndex*32 : (entityIndex+1)*32]
  p.Memories[21] = p.Memories[11][entityIndex*32 : (entityIndex+1)*32]
  p.Memories[22] = p.Memories[12][entityIndex*32 : (entityIndex+1)*32]
	switch status {
	case StatusPreprocess:
		p.MemoryPermissions = map[int]MemoryPermission{
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
		p.MemoryPermissions = map[int]MemoryPermission{
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
		p.MemoryPermissions = map[int]MemoryPermission{
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
		p.MemoryPermissions = map[int]MemoryPermission{
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
		p.MemoryPermissions = map[int]MemoryPermission{
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
		p.MemoryPermissions = map[int]MemoryPermission{
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
		p.MemoryPermissions = map[int]MemoryPermission{
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
		p.MemoryPermissions = map[int]MemoryPermission{
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
