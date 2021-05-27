package bridging

import "fmt"

type HybridEngine struct {
	Engine
}

func (h *HybridEngine) Start() {
	fmt.Println("HybridEngine")
}
