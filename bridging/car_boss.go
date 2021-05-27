package bridging

type BossCar struct {
	RefinedCar
}

func NewBossCar(engine Engine) *BossCar {
	return &BossCar{NewRefinedCar(engine)}
}

func (b *BossCar) GetBrand() string {
	return "Boss"
}
