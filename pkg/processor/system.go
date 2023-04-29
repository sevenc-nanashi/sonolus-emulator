package processor

func (p Processor) Prepare() {
	for i, entity := range p.Entities {
    entity.Archetype.Preprocess(i)
	}
}
