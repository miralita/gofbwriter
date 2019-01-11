package gofbwriter

//ShareInstruction - In-document instruction for generating output free and payed documents
type ShareInstruction struct {
	b       *builder
	tagName string
}

func (s *ShareInstruction) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//ToXML - export to XML string
func (*ShareInstruction) ToXML() (string, error) {
	panic("implement me")
}

func (s *ShareInstruction) tag() string {
	if s.tagName == "" {
		s.tagName = "share-instruction-type"
	}
	return s.tagName
}
