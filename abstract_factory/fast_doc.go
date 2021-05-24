package abstract_factory

type (
	FastFactory      struct{}
	FastHtmlDocument struct {
		Markdown string
	}
	FastWordDocument struct {
		Markdown string
	}
)

func (f FastFactory) CreateHTML(markdown string) HTMLDocument {
	return &FastHtmlDocument{Markdown: markdown}
}

func (f FastFactory) CreateWord(markdown string) WordDocument {
	return &FastWordDocument{Markdown: markdown}
}

func (f *FastHtmlDocument) ToHTML() string {
	return ""
}

func (f *FastHtmlDocument) Save(path string) error {
	return nil
}

func (f *FastWordDocument) Save(path string) error {
	return nil
}
