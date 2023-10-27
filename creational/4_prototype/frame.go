package prototype

type Frame interface {
	OpenTheDoor()
	CloseTheDoor()
	Clone() Frame
}

type frame struct {
	doors [2]Door
	//windows [3]Window
}

func (f *frame) OpenTheDoor() {
	for i := 0; i < len(f.doors); i++ {
		f.doors[i].Open()
	}
}

func (f *frame) CloseTheDoor() {
	for i := 0; i < len(f.doors); i++ {
		f.doors[i].Close()
	}
}

func (f *frame) Clone() Frame {
	nf := new(frame)
	nf.doors[0] = f.doors[0].Clone()
	nf.doors[1] = f.doors[1].Clone()
	return nf
}

type Door interface {
	Open()
	Close()
	Clone() Door
}

type door struct {
	vetexes []vertex
}

func (d *door) Open() {
}

func (d *door) Close() {
}

func (d *door) Clone() Door {
	nd := new(door)
	vv := make([]vertex, len(d.vetexes))
	vv = append(vv[0:0], d.vetexes...)
	nd.vetexes = vv
	return nd
}

type vertex struct {
	x, y      float32
	thickness float32
}
