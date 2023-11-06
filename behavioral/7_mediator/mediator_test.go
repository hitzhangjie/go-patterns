package mediator

import (
	"testing"
	"time"
)

func TestMediator(t *testing.T) {
	mediator := NewFittingRoomMediator(2)

	c1 := NewCustomer("zhang").SetMediator(mediator)
	c2 := NewCustomer("yang").SetMediator(mediator)
	c3 := NewCustomer("wang").SetMediator(mediator)

	ch := make(chan int, 2)
	go func() {
		c1.Enter()
		c2.Enter()
		c3.Enter()
		ch <- 1
	}()
	go func() {
		time.Sleep(time.Second * 5)
		c1.Leave()
		c2.Leave()
		time.Sleep(time.Second * 5)
		c3.Leave()
		ch <- 1
	}()
	<-ch
	<-ch
}
