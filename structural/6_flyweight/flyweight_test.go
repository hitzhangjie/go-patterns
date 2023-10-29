package flyweight_test

import (
	"flyweight"
	"testing"
)

func TestFlyweight(t *testing.T) {
	c := flyweight.NewCircle(10, 10, 5, flyweight.Red, 0.3)
	c.Draw(flyweight.Pos{10, 10})
}
