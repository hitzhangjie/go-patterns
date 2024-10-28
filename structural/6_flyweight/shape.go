package flyweight

import (
	"fmt"
	"math"
)

// Shape draw this shape at pos
type Shape interface {
	Draw()
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
// Here we doesn't consider safe when use it concurrently.
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
	Pos        Pos
	LineWeight float64
}

// Draw draws the point at specific position using specific color and line
// weight.
//
// Note: I'm not a professional Computer Graphics engineers, I really don't
// know what's the differences btw a point or a particle, here I just try
// to treat a Point as a thing which may be rendered in many Particles.
func (p *Point) Draw() {
	fmt.Printf("\\--> draw point at <%.1f, %.1f> with color:%s with weight:%.1f\n",
		p.Pos.X,
		p.Pos.Y,
		p.color,
		p.LineWeight)
	for i := 0.0; i < p.LineWeight; i += 0.1 {
		for j := 0.0; j < p.LineWeight; j += 0.1 {
			p.Particle.Draw(Pos{p.Pos.X + i, p.Pos.Y + j})
		}
	}
}

func NewCircle(pos Pos, raidus float64, color Color, weight float64) *Circle {
	return &Circle{
		Color:      color,
		Center:     pos,
		Radius:     raidus,
		LineWeight: weight,
	}
}

// Circle a circle
type Circle struct {
	Color      Color
	Center     Pos
	Radius     float64
	LineWeight float64
}

// Draw circle will be rendered by different Points according to the formula
// (x-c.X)^2 + (y-c.Y)^2 = c.Raidus^2
func (c Circle) Draw() {
	fmt.Printf("draw circle at <%.1f, %.1f> with raidus %.1fcm, weight:%.1f\n",
		c.Center.X,
		c.Center.Y,
		c.Radius,
		c.LineWeight)

	// (x-pos.X)^2 + (y-pos.Y)^2 = c.Radius^2
	for x := c.Center.X - c.Radius; x < c.Center.X+c.Radius; x += 0.1 {
		y := math.Pow(math.Pow(c.Radius, 2)-math.Pow(x-c.Center.X, 2), 0.5) + c.Center.Y
		p := &Point{
			Particle:   pf.GetParticle(c.Color),
			LineWeight: c.LineWeight,
		}
		p.Pos = Pos{x, y}
		p.Draw()

		p.Pos = Pos{x, -y}
		p.Draw()
	}
}

type Pos struct {
	X, Y float64
}
