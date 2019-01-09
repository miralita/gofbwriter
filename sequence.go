package gofbwriter

type sequence struct {
	b *builder
}

func (s *sequence) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (*sequence) ToXML() (string, error) {
	panic("implement me")
}

func (s *sequence) tag() string {
	return "sequence"
}
