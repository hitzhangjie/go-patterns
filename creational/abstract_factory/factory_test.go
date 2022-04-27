package abstract_factory_test

import (
	"testing"

	af "github.com/hitzhangjie/go-patterns/creational/abstract_factory"
)

func Test_Creational_AbstractFactory(t *testing.T) {
	factory := af.Get(af.Style1)
	factory.CreateProductA().DoSomethingA()
	factory.CreateProductB().DoSomethingB()

	factory = af.Get(af.Style2)
	factory.CreateProductA().DoSomethingA()
	factory.CreateProductB().DoSomethingB()
}
