package gofbwriter

type section struct {
	b        *builder
	title    *title
	epigraph *epigraph
	image    *image
}

func (s *section) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (*section) ToXML() (string, error) {
	panic("implement me")
}

func (s *section) tag() string {
	return "section"
}
