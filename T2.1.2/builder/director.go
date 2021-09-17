package builder

type directorInterface interface {
	setBuilder(b iBuilder)
	buildHouse() house
}

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) directorInterface {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
