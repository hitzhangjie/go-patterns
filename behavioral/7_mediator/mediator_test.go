package mediator

import "testing"

func TestMediator(t *testing.T) {
	c1 := NewCustomer("zhang")
	c2 := NewCustomer("yang")
	c3 := NewCustomer("wang")

	mediator := NewFittingRoomMediator(2)
	c1.Enter(mediator)
	c2.Enter(mediator)
	c3.Enter(mediator)
	c1.Leave(mediator)
	c2.Leave(mediator)
	c3.Leave(mediator)
}
