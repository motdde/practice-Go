//Package shapes is to demonstrate the implementation of interfaces
package shapes

import (
	"math"
)

//Geometry interface for shapes:
type Geometry interface {
	Area() float64
	Perim() float64
}

//Rect type
type Rect struct {
	Width, Height float64
}

//Circle type
type Circle struct {
	Radius float64
}

// Area calulates the area
func (r *Rect) Area() float64 {
	return r.Width * r.Height
}

// Perim calulates the permieter
func (r *Rect) Perim() float64 {
	return 2*r.Width + 2*r.Height
}

// Area calulates the area
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perim calulates the permieter
func (c *Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}
