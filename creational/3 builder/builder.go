package builder

type builder struct {
	house House
}

func NewBuilder() *builder {
	return &builder{}
}

func (b *builder) BuildRooves(style RoofStyle) *builder {
	// build the rooves according to the walls' info
	return b
}

func (b *builder) BuildWall(pos Pos, color ColorStyle) *Wall {
	return &Wall{
		Pos:   pos,
		Color: color,
	}
}

func (w *Wall) BuildWindow(pos Pos, height, width int, style GlassStyle) *Wall {
	w.Window = append(w.Window, Window{
		Pos:    pos,
		Height: height,
		Width:  width,
		Glass:  style,
	})
	return w
}

func (w *Wall) BuildDoor(pos Pos, height, width int, style MaterialStyle, color ColorStyle) *Wall {
	w.Door = append(w.Door, Door{
		Pos:      pos,
		Material: style,
		Height:   height,
		Width:    width,
		Color:    color,
	})
	return w
}

func (b *builder) BuildFloor(pos []Pos, style MaterialStyle) *builder {
	b.house.Floor = append(b.house.Floor, Floor{
		Vetexes:  pos,
		Material: style,
	})
	return b
}

func (b *builder) Build() House {
	return b.house
}
