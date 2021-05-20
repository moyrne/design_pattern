package factory

import "fmt"

type Shape interface {
	Print()
}

type Slice []interface{}

func (s Slice) Print() {
	for _, v := range s {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
}

type Link struct {
	Value interface{}
	Next  *Link
}

func (l *Link) Print() {
	p := l
	for p != nil {
		fmt.Printf("%v\t", p.Value)
		p = p.Next
	}
	fmt.Println()
}

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "slice":
		return Slice{1, 2, 3, 4, 5, 6}
	case "link":
		return &Link{1, &Link{2, &Link{3, &Link{4, &Link{5, &Link{Value: 6}}}}}}
	default:
		return nil
	}
}
