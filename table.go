package gofbwriter

//Basic html-like tables
type table struct {
	b *builder
}

func (s *table) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *table) ToXML() (string, error) {
	panic("implement me")
}

func (s *table) tag() string {
	return "table"
}
