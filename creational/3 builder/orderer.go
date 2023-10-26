package builder

type Orderer struct {
	b       *builder
	Floors  []Floor
	Wall    []Wall
	Windows []Window
	Doors   []Door
}

func NewOrderer(b *builder) *Orderer {
	return &Orderer{b: b}
}

func (ord *Orderer) WantFloor(vetexes []Pos, style MaterialStyle) {
	ord.Floors = append(ord.Floors, Floor{
		Vetexes:  vetexes,
		Material: style,
	})
}

func (ord *Orderer) WantWindow(pos Pos, height, width int, style GlassStyle) {
	ord.Windows = append(ord.Windows, Window{
		Pos:    pos,
		Height: height,
		Width:  width,
		Glass:  style,
	})
}

func (ord *Orderer) WantWall(pos Pos, color ColorStyle) {
	ord.Wall = append(ord.Wall, Wall{
		Pos:    pos,
		Window: nil,
		Door:   nil,
		Color:  0,
	})
}

func (ord *Orderer) Build() House {
	// prepare the building elements, like which windows and doors belongs to the 1st wall, 2nd wall, etc.
	// ...

	// then:
	// - call ord.b.BuildFloor
	// - call ord.b.BuildWall with ord.b.BuildWindow and ord.b.BuildDoor
	// - call ord.b.BuildRooves
	// - call ord.b.Build() to get the final House
	return ord.b.Build()
}
