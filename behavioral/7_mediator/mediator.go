package mediator

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Mediator defines the fitting room mediator behavior
type Mediator interface {
	// Wait let customer wait untils a fitting room released
	Wait(*Customer) <-chan *FittingRoom

	// Enter let customer enter a fitting room
	Enter(*Customer)

	// Leave let customer leave a fitting room
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

// a customer who waits for a fitting room
type waiter struct {
	*Customer
	ch chan *FittingRoom
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
	log.Printf("%s no reserved room found\n", customer.Name)
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
			log.Printf("%s will leave room-%d\n", customer.Name, room.no)
			// let next waiting customer enter
			select {
			case waiter := <-m.WaitingCustomers:
				log.Printf("%s will be notified to use room-%d\n", waiter.Name, room.no)
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
		log.Printf("%s wait too long, it's best to quit\n", c.Name)
		return
	}
	log.Printf("%s could use reserved room-%d\n", c.Name, room.no)

	c.mediator.Enter(c)
	log.Printf("%s now is using room-%d\n", c.Name, room.no)
}

func (c *Customer) Leave() {
	c.mediator.Leave(c)
	log.Printf("%s has left\n", c.Name)
}
