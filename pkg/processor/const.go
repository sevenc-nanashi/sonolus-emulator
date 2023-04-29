package processor

type Status string

const (
	StatusPreprocess       = Status("Preprocess")
	StatusSpawnOrder       = Status("SpawnOrder")
	StatusShouldSpawn      = Status("ShouldSpawn")
	StatusInitialize       = Status("Initialize")
	StatusUpdateSequential = Status("UpdateSequential")
	StatusTouch            = Status("Touch")
	StatusUpdateParallel   = Status("UpdateParallel")
	StatusTerminate        = Status("Terminate")
)
