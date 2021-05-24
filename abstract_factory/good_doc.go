package abstract_factory

type (
	GoodFactory      struct{}
	GoodHTMLDocument struct {
		Markdown string
	}
	GoodWordDocument struct {
		Markdown string
	}
)

func (g GoodFactory) CreateHTML(markdown string) HTMLDocument {
	return &GoodHTMLDocument{Markdown: markdown}
}

func (g GoodFactory) CreateWord(markdown string) WordDocument {
	return &GoodWordDocument{Markdown: markdown}
}

func (g GoodHTMLDocument) ToHTML() string {
	return ""
}

func (g GoodHTMLDocument) Save(path string) error {
	return nil
}

func (g GoodWordDocument) Save(path string) error {
	return nil
}
