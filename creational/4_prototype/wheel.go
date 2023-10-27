package prototype

type Wheel interface {
	Spin()
	Brake()
	Clone() Wheel
}

type wheel struct {
	tire *tire
	tube *tube
}

type tire struct{}
type tube struct{}

func (w *wheel) Spin() {
}

func (w *wheel) Brake() {
}

func (w *wheel) Clone() Wheel {
	nw := new(wheel)
	nw.tire = new(tire)
	nw.tube = new(tube)
	return nw
}
