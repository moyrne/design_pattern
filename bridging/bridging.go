package bridging

import "fmt"

type (
	Engine interface {
		Start()
	}

	Car struct {
		Engine
	}
)

func (c *Car) Drive() {

}

type RefinedCar struct {
	Car
}

func NewRefinedCar(engine Engine) RefinedCar {
	return RefinedCar{Car{Engine: engine}}
}

func (r *RefinedCar) Drive() {
	r.Engine.Start()
}

func (r *RefinedCar) GetBrand() string {
	return ""
}

type BossCar struct {
	RefinedCar
}

func NewBossCar(engine Engine) *BossCar {
	return &BossCar{NewRefinedCar(engine)}
}

func (b *BossCar) GetBrand() string {
	return "Boss"
}

type HybridEngine struct {
	Engine
}

func (h *HybridEngine) Start() {
	fmt.Println("HybridEngine")
}
