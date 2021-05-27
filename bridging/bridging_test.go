package bridging

import "testing"

func TestNewBossCar(t *testing.T) {
	car := NewBossCar(&HybridEngine{})
	car.Drive()
}
