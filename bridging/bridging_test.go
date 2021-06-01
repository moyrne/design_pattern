package bridging

import "testing"

// 这是桥接模式?
func TestNewBossCar(t *testing.T) {
	bossHybridCar := NewBossCar(NewHybridEngine())
	bossHybridCar.Drive()
	refinedHybridCar := NewRefinedCar(NewHybridEngine())
	refinedHybridCar.Drive()

	bossTractorCar := NewBossCar(NewTractorEngine())
	bossTractorCar.Drive()
	refinedTractorCar := NewRefinedCar(NewTractorEngine())
	refinedTractorCar.Drive()
}
