package bridging

import "fmt"

type HybridEngine struct{}

func NewHybridEngine() *HybridEngine {
	return &HybridEngine{}
}

func (h *HybridEngine) Start() {
	fmt.Println("Start Hybrid Engine")
}

type TractorEngine struct{}

func NewTractorEngine() *TractorEngine {
	return &TractorEngine{}
}

func (t *TractorEngine) Start() {
	fmt.Println("Start Tractor Engine")
}
