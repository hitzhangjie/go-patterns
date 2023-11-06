package mediator

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Mediator interface {
	Wait(*Customer) <-chan *FittingRoom
	Enter(*Customer)
	Leave(*Customer)
}

const (
	idle = iota
	reserved
	inused
)

type FittingRoom struct {
	no       int
	state    int
	customer *Customer
	last     time.Time
}

// NewFittingRoomMediator create a mediator who manages `cap` fitting rooms
func NewFittingRoomMediator(cap int) *fittingRoomMediator {
	rooms := make([]*FittingRoom, cap)
	for i := 0; i < cap; i++ {
		rooms[i] = &FittingRoom{
			no:       i + 1,
			state:    idle,
			customer: nil,
		}
	}
	m := &fittingRoomMediator{
		Rooms:            rooms,
		WaitingCustomers: make(chan *waiter, 64),
	}
	return m
}

type fittingRoomMediator struct {
	sync.Mutex
	Rooms []*FittingRoom

	WaitingCustomers chan *waiter
}

type waiter struct {
	*Customer
	ch chan *FittingRoom
}

func (m *fittingRoomMediator) Wait(customer *Customer) <-chan *FittingRoom {
	m.Lock()
	defer m.Unlock()

	ch := make(chan *FittingRoom, 1)
	for _, room := range m.Rooms {
		if room.state == idle ||
			(room.state == reserved && time.Since(room.last) > time.Minute*5) {
			room.state = reserved
			room.customer = customer
			room.last = time.Now()
			ch <- room
			return ch
		}
	}

	m.WaitingCustomers <- &waiter{Customer: customer, ch: ch}
	return ch
}

func (m *fittingRoomMediator) Enter(customer *Customer) {
	m.Lock()
	defer m.Unlock()

	for _, room := range m.Rooms {
		if room.state != reserved {
			continue
		}
		if room.customer.sequence == customer.sequence {
			room.state = inused
			room.last = time.Now()
			return
		}
	}
	fmt.Printf("%s no reserved room found\n", customer.Name)
}

func (m *fittingRoomMediator) Leave(customer *Customer) {
	m.Lock()
	defer m.Unlock()

	for _, room := range m.Rooms {
		if room.state != inused {
			continue
		}
		if room.customer.sequence == customer.sequence {
			room.state = idle
			room.customer = nil
			room.last = time.Now()
			fmt.Printf("%s will leave room-%d\n", customer.Name, room.no)
			// let next waiting customer enter
			select {
			case waiter := <-m.WaitingCustomers:
				fmt.Printf("%s will be notified to use room-%d\n", waiter.Name, room.no)
				room.state = reserved
				room.customer = waiter.Customer
				room.last = time.Now()
				waiter.ch <- room
				close(waiter.ch)
			default:
			}
		}
	}
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
	mediator Mediator
}

func (c *Customer) SetMediator(m Mediator) *Customer {
	c.mediator = m
	return c
}

func (c *Customer) Enter() {
	room, ok := <-c.mediator.Wait(c)
	if !ok {
		fmt.Printf("%s wait too long, it's best to quit\n", c.Name)
		return
	}
	fmt.Printf("%s could use reserved room-%d\n", c.Name, room.no)

	c.mediator.Enter(c)
	fmt.Printf("%s now is using room-%d\n", c.Name, room.no)
}

func (c *Customer) Leave() {
	c.mediator.Leave(c)
	fmt.Printf("%s has left\n", c.Name)
}
