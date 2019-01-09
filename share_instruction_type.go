package gofbwriter

//In-document instruction for generating output free and payed documents
type shareInstructionType struct {
	b       *builder
	tagName string
}

func (s *shareInstructionType) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (*shareInstructionType) ToXML() (string, error) {
	panic("implement me")
}

func (s *shareInstructionType) tag() string {
	if s.tagName == "" {
		s.tagName = "share-instruction-type"
	}
	return s.tagName
}
