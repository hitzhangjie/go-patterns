package mediator

import (
	"fmt"
	"sync/atomic"
)

type Mediator interface {
	Enter(*Customer) bool
	Leave(*Customer) bool
}

type fittingRoom struct {
	no       int
	inUse    bool
	customer *Customer
}

func NewFittingRoomMediator(numOfRooms int) *fittingRoomMediator {
	rooms := make([]*fittingRoom, numOfRooms)
	for i := 0; i < numOfRooms; i++ {
		rooms[i] = &fittingRoom{
			no:       i + 1,
			inUse:    false,
			customer: nil,
		}
	}
	m := &fittingRoomMediator{
		Rooms:            rooms,
		WaitingCustomers: make(chan *Customer, 64),
	}
	return m
}

type fittingRoomMediator struct {
	Rooms            []*fittingRoom
	WaitingCustomers chan *Customer
}

func (m *fittingRoomMediator) Enter(customer *Customer) bool {
	for _, room := range m.Rooms {
		if !room.inUse {
			room.inUse = true
			room.customer = customer
			fmt.Printf("%s enter room-%d\n", customer.Name, room.no)
			return true
		}
	}
	fmt.Printf("%s wait idle rooms\n", customer.Name)
	m.WaitingCustomers <- customer
	return false
}

func (m *fittingRoomMediator) Leave(customer *Customer) bool {
	for _, room := range m.Rooms {
		if room.customer == nil {
			continue
		}
		if room.customer.sequence == customer.sequence {
			room.inUse = false
			room.customer = nil
			fmt.Printf("%s leave room-%d\n", customer.Name, room.no)
			// let next waiting customer enter
			select {
			case nextCustomer := <-m.WaitingCustomers:
				m.Enter(nextCustomer)
			default:
			}
		}
	}
	return true
}

var sequence atomic.Int64

func NewCustomer(name string) *Customer {
	return &Customer{
		Name:     name,
		sequence: int(sequence.Add(1)),
	}
}

type Customer struct {
	Name     string
	sequence int
}

func (c *Customer) Enter(mediator Mediator) {
	mediator.Enter(c)
}

func (c *Customer) Leave(mediator Mediator) {
	mediator.Leave(c)
}
