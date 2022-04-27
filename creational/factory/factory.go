package factory

type IFactory interface {
	CreateProductA() IProductA
	CreateProductB() IProductB
}

type Factory struct{}

func NewFactory() IFactory {
	return &Factory{}
}

func (f Factory) CreateProductA() IProductA {
	return &ProductA{}
}

func (f Factory) CreateProductB() IProductB {
	return &ProductB{}
}
