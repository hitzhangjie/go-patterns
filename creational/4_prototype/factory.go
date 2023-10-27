package prototype

// interface{} is discovered not by designed in advanced, here
// `type Factory struct{}` is ok.
//
// if we will create different kinds of cars, we may need different factories,
// then we'll need an `type Factory interface{}` to specify the behavior of a
// factory, and add subclasses to implement this interface.
type Factory struct{}

func (f *Factory) Create() Car {
	c := new(car)

	c.Frame = &frame{
		doors: [2]Door{
			&door{vetexes: []vertex{{0, 0, 5}, {0, 100, 5}, {100, 100, 5}, {0, 100, 5}}},
			&door{vetexes: []vertex{{0, 0, 5}, {0, 100, 5}, {100, 100, 5}, {0, 100, 5}}},
		},
	}

	for i := 0; i < len(c.Wheels); i++ {
		c.Wheels[i] = &wheel{
			tire: &tire{},
			tube: &tube{},
		}
	}

	//...

	return c
}
