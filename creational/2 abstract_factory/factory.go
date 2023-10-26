package abstract_factory

// IFactory defines factory's behaviour
type IFactory interface {
	Style() Style
	CreateProductA() IProductA
	CreateProductB() IProductB
}

// Style1Factory create products with style1
type Style1Factory struct{}

func (f Style1Factory) Style() Style {
	return Style1
}

func (f Style1Factory) CreateProductA() IProductA {
	return &Style1ProductA{}
}

func (f Style1Factory) CreateProductB() IProductB {
	return &Style1ProductB{}
}

// Style2Factory create products with style2
type Style2Factory struct{}

func (f Style2Factory) Style() Style {
	return Style2
}

func (f Style2Factory) CreateProductA() IProductA {
	return &Style2ProductA{}
}

func (f Style2Factory) CreateProductB() IProductB {
	return &Style2ProductB{}
}
