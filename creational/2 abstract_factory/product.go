package abstract_factory

// IProductA defines behavior of ProductA
type IProductA interface {
	DoSomethingA() error
}

// IProductB defines behavior of ProductB
type IProductB interface {
	DoSomethingB() error
}

type Style1ProductA struct{}

func (p Style1ProductA) DoSomethingA() error {
	println("i am product-a with style1")
	return nil
}

type Style2ProductA struct{}

func (p Style2ProductA) DoSomethingA() error {
	println("i am product-a with style2")
	return nil
}

type Style1ProductB struct{}

func (p Style1ProductB) DoSomethingB() error {
	println("i am product-b with style1")
	return nil
}

type Style2ProductB struct{}

func (p Style2ProductB) DoSomethingB() error {
	println("i am product-b with style2")
	return nil
}
