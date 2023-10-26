package factory_test

import (
	"testing"

	"github.com/hitzhangjie/go-patterns/creational/factory"
)

func Test_Creational_Factory(t *testing.T) {
	f := factory.NewFactory()
	if err := f.CreateProductA().DoSomethingA(); err != nil {
		t.Errorf("create ProductA error: %v", err)
	}
	if err := f.CreateProductB().DoSomethingB(); err != nil {
		t.Errorf("create ProductB error: %v", err)
	}
}
