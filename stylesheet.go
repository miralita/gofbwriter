package gofbwriter

type stylesheet struct {
	b     *builder
	ctype string
	data  string
	book  *book
}

func (s *stylesheet) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *stylesheet) Set(ctype, data string) {
	s.ctype = ctype
	s.data = data
}

func (s *stylesheet) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
	b.openTagAttr(s.tag(), map[string]string{"type": s.ctype}, false)
	b.WriteString(s.data)
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *stylesheet) tag() string {
	return "stylesheet"
}
