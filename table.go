package gofbwriter

//Table - basic html-like tables
type Table struct {
	b *builder
}

func (s *Table) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//ToXML - export to XML string
func (s *Table) ToXML() (string, error) {
	panic("implement me")
}

func (s *Table) tag() string {
	return "table"
}
