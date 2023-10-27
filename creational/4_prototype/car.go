package prototype

type Direction int

const (
	SteerLeft Direction = iota
	SteerRight
)

type Car interface {
	Accelerate()
	Brake()
	Steer(Direction)
	Clone() Car
}

type car struct {
	Frame  Frame
	Wheels [4]Wheel
	//Engine
	//Lights
	//FuelTank
	//...
}

func (c *car) Accelerate() {
	panic("not implemented")
}

func (c *car) Brake() {
	panic("not implemented")
}

func (c *car) Steer(_ Direction) {
	panic("not implemented")
}

func (c *car) Clone() Car {
	nc := new(car)
	// clone frame
	nc.Frame = c.Frame.Clone()
	// clone wheels
	nc.Wheels[0] = c.Wheels[0].Clone()
	nc.Wheels[1] = c.Wheels[1].Clone()
	nc.Wheels[2] = c.Wheels[2].Clone()
	nc.Wheels[3] = c.Wheels[3].Clone()
	// clone others like engine, fuel tank, lighs
	// ...
	return nc
}
