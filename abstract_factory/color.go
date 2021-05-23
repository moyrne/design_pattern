package abstract_factory

import "fmt"

type (
	ColorFactory struct{}
	Color        interface {
		Fill()
	}
	RGB string
)

func (c *ColorFactory) GetColor(color string) Color {
	switch color {
	case "red":
		return &Red{RGB: "Red"}
	case "green":
		return &Green{RGB: "Green"}
	case "blue":
		return &Blue{RGB: "Blue"}
	}
	return nil
}

type (
	Red struct {
		RGB
	}
	Green struct {
		RGB
	}
	Blue struct {
		RGB
	}
)

func (r *Red) Fill() {
	fmt.Println("red")
}

func (g *Green) Fill() {
	fmt.Println("green")
}

func (b *Blue) Fill() {
	fmt.Println("blue")
}
