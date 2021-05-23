package abstract_factory

import (
	"fmt"
	"math"
)

type (
	ShapeFactory struct{}
	Shape        interface {
		Draw()
	}
	Circle struct {
		R int
	}
	Square struct {
		W, H int
	}
	Rectangle struct {
		B, H int
	}
)

func (s *ShapeFactory) GetShape(shape string) Shape {
	switch shape {
	case "circle":
		return &Circle{R: 1}
	case "square":
		return &Square{W: 1, H: 1}
	case "rectangle":
		return &Rectangle{B: 1, H: 1}
	default:
		return nil
	}
}

func (c *Circle) Draw() {
	fmt.Println("circle", float64(c.R)*math.Pi)
}

func (s *Square) Draw() {
	fmt.Println("square", s.W*s.H)
}

func (r *Rectangle) Draw() {
	fmt.Println("rectangle", r.B*r.H/2)
}
