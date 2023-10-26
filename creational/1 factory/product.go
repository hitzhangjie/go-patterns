package factory

type IProductA interface {
	DoSomethingA() error
}

type IProductB interface {
	DoSomethingB() error
}

type ProductA struct{}

func (p ProductA) DoSomethingA() error {
	return nil
}

type ProductB struct{}

func (p ProductB) DoSomethingB() error {
	return nil
}
