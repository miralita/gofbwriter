package gofbwriter

type annotation struct {
	b       *builder
	tagName string
}

func (s *annotation) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (*annotation) ToXML() (string, error) {
	panic("implement me")
}

func (s *annotation) AddParagraph(p string) {
	panic("implement me")
}

func (s *annotation) tag() string {
	if s.tagName == "" {
		s.tagName = "annotation"
	}
	return s.tagName
}
