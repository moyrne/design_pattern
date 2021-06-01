package bridging

import "fmt"

type BossCar struct {
	engine Engine
}

func NewBossCar(engine Engine) *BossCar {
	return &BossCar{engine}
}

func (b *BossCar) Drive() {
	b.engine.Start()
	fmt.Println("Drive Boss Car")
	fmt.Println("--------------")
}
