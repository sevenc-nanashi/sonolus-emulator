package processor

func (p *Processor) Prepare() {
	for i, entity := range p.Entities {
		entity.Archetype.Preprocess(i)
	}
	for i, entity := range p.Entities {
		entity.SpawnOrder = entity.Archetype.SpawnOrder(i)
	}
}

func (p *Processor) Update(currentTime float64, deltaTime float64) {
	p.memories[1][0] = currentTime
	p.memories[1][1] = deltaTime
	for i, entity := range p.Entities {
		if entity.initialized {
			continue
		}
		entity.Archetype.Initialize(i)
		p.Entities[i].initialized = true
	}
	for i, entity := range p.Entities {
		entity.Archetype.UpdateSequential(i)
	}
}

type TouchInfo struct {
	Id        int
	Status    TouchStatus
	Time      float64
	StartTime float64
	X         float64
	Y         float64
	Sx        float64
	Sy        float64
	Dx        float64
	Dy        float64
}

type TouchStatus string

const (
	TouchStatusStart  TouchStatus = "start"
	TouchStatusMiddle TouchStatus = "middle"
	TouchStatusEnd    TouchStatus = "end"
)

func (p *Processor) Touch(touches []TouchInfo) {
	for _, touch := range touches {
		touchData := []float64{
			float64(touch.Id),
			-1,
			-1,
			touch.Time,
			touch.StartTime,
			touch.X,
			touch.Y,
			touch.Sx,
			touch.Sy,
			touch.Dx,
			touch.Dy,
			0, // TODO: vx
			0, // TODO: vy
			0, // TODO: vr
			0, // TODO: vw
		}

		switch touch.Status {
		case TouchStatusStart:
			touchData[1] = 1
			touchData[2] = 0
			break
		case TouchStatusMiddle:
			touchData[1] = 0
			touchData[2] = 0
			break
		case TouchStatusEnd:
			touchData[1] = 0
			touchData[2] = 1
			break
		}

		p.memories[101] = touchData

		for i, entity := range p.Entities {
			entity.Archetype.Touch(i)
		}
	}
}

func (p *Processor) Spawn() {
	for _, entity := range p.spawnQueue {
		p.Entities = append(p.Entities, Entity{
			Archetype: Archetype{
				Processor: p,
				Index:     -1,
				Scripts:   entity.Scripts,
			},
			SpawnOrder: -1,
			Values:     []float64{},
		})
		memory := make([]float64, 32)
		for i, value := range entity.Data {
			memory[i] = value
		}
		p.EntityMemories[len(p.Entities)-1] = EntityMemories{
			Memory: memory,
			Input:  make([]float64, 4),
		}

	}
	p.spawnQueue = make([]SpawnQueueItem, 0)
}
