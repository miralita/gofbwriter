package gofbwriter

type section struct {
	title    *title
	epigraph *epigraph
	image    *image
}

func (*section) ToXML() (string, error) {
	panic("implement me")
}

func (s *section) tag() string {
	return "section"
}
