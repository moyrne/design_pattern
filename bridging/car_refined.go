package bridging

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
