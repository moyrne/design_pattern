package bridging

import "fmt"

type RefinedCar struct {
	engine Engine
}

func NewRefinedCar(engine Engine) RefinedCar {
	return RefinedCar{engine}
}

func (r *RefinedCar) Drive() {
	r.engine.Start()
	fmt.Println("Drive Refined Car")
	fmt.Println("-----------------")
}
