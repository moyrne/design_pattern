package abstract_factory

type AbstractFactory interface {
	GetColor() Color
	GetShape() Shape
}
