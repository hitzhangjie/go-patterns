package flyweight

import (
	"fmt"
	"math"
)

type Shape interface {
	Draw(pos Pos)
}

type Color string

const (
	Red   = "red"
	Blue  = "blue"
	Green = "green"
)

var pf = &ParticleFactory{
	particles: make(map[Color]*Particle),
}

// ParticleFactory creates particles according to `Color`
type ParticleFactory struct {
	particles map[Color]*Particle
}

// GetParticle returns reusable underlying particle instance with `Color`
func (pf ParticleFactory) GetParticle(c Color) *Particle {
	if v, ok := pf.particles[c]; ok {
		return v
	}
	v := &Particle{c}
	pf.particles[c] = v
	return v
}

// Particle maintains the intrinsic state including color
type Particle struct {
	color Color
	// ...
	// ...
}

// Draw draw particle at `pos`, pos is the extrinsic state maintained
// by Point or Circle
func (c *Particle) Draw(pos Pos) {
	fmt.Printf("\t\\--> draw particle at <%.1f, %.1f> with color:%s\n",
		pos.X,
		pos.Y,
		c.color)
}

// Point a point which may be moved to <x,y> and rendered by different
// line weight
type Point struct {
	*Particle
	X, Y       float64
	LineWeight float64
}

// Draw draws the point at specific position using specific color and line
// weight.
//
// Note: I'm not a professional Computer Graphics engineers, I really don't
// know what's the differences btw a point or a particle, here I just try
// to treat a Point as a thing which may be rendered in many Particles.
func (p *Point) Draw(pos Pos) {
	fmt.Printf("\\--> draw point at <%.1f, %.1f> with color:%s with weight:%.1f\n",
		pos.X,
		pos.Y,
		p.color,
		p.LineWeight)
	for i := 0.0; i < p.LineWeight; i += 0.1 {
		p.Particle.Draw(Pos{pos.X + i, pos.Y + i})
	}
}

func NewCircle(x, y, raidus float64, color Color, weight float64) *Circle {
	return &Circle{
		Color:      color,
		Radius:     raidus,
		LineWeight: weight,
	}
}

// Circle a circle
type Circle struct {
	Color      Color
	Radius     float64
	LineWeight float64
}

// Draw circle will be rendered by different Points according to the formula
// (x-c.X)^2 + (y-c.Y)^2 = c.Raidus^2
func (c Circle) Draw(pos Pos) {
	fmt.Printf("draw circle at <%.1f, %.1f> with raidus %.1fcm, weight:%.1f\n",
		pos.X,
		pos.Y,
		c.Radius,
		c.LineWeight)
	// (x-pos.X)^2 + (y-pos.Y)^2 = c.Radius^2
	for x := pos.X - c.Radius; x < pos.X+c.Radius; x += 0.1 {
		y := math.Pow(math.Pow(c.Radius, 2)-math.Pow(x-pos.X, 2), 0.5) + pos.Y
		p := &Point{
			Particle:   pf.GetParticle(c.Color),
			X:          x,
			Y:          y,
			LineWeight: c.LineWeight,
		}
		p.Draw(Pos{p.X, p.Y})
	}
}

type Pos struct {
	X, Y float64
}
