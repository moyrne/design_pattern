package bridging

import "fmt"

type Car struct {
	Engine
}

func (c *Car) Drive() {
	fmt.Println("Car.Drive")
}
