package builder

type House struct {
	Floor  []Floor
	Walls  []Wall
	Rooves []Roof
}

type Wall struct {
	Pos
	Window []Window
	Door   []Door
	Color  ColorStyle
}

type Floor struct {
	Vetexes  []Pos
	Material MaterialStyle
}

type Roof struct {
	Style RoofStyle
}

type Window struct {
	Pos
	Height int
	Width  int
	Glass  GlassStyle
}

type Door struct {
	Pos
	Material MaterialStyle
	Height   int
	Width    int
	Color    ColorStyle
}

type MaterialStyle int
type ColorStyle int
type GlassStyle int
type RoofStyle int

type Pos struct {
	X int
	Y int
	Z int
}
