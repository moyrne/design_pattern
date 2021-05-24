package abstract_factory

import "github.com/pkg/errors"

// https://www.liaoxuefeng.com/wiki/1252599548343744/1281319134822433

type (
	AbstractFactory interface {
		CreateHTML(markdown string) HTMLDocument
		CreateWord(markdown string) WordDocument
	}
	HTMLDocument interface {
		ToHTML() string
		Save(path string) error
	}
	WordDocument interface {
		Save(path string) error
	}
)

var ErrNoFactory = errors.New("not found factory")

func CreateFactory(factory string) (AbstractFactory, error) {
	switch factory {
	case "fast":
		return FastFactory{}, nil
	case "good":
		return GoodFactory{}, nil
	default:
		return nil, errors.WithMessage(ErrNoFactory, factory)
	}
}

func DoSomeThing(factory string) {
	fastFactory, err := CreateFactory(factory)
	if err != nil {
		return
	}
	html := fastFactory.CreateHTML("")
	if err := html.Save(""); err != nil {
		return
	}

	word := fastFactory.CreateWord("")
	if err := word.Save(""); err != nil {
		return
	}
}
