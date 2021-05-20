package factory

import (
	"testing"
)

func TestShapeFactory(t *testing.T) {
	slice := ShapeFactory("slice")
	slice.Print()

	link := ShapeFactory("link")
	link.Print()
}
