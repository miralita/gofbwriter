package gofbwriter

//CustomInfo - Any other information about the book/document
type CustomInfo struct {
	b        *builder
	info     string
	infoType string
}

func (s *CustomInfo) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//ToXML - export to XML string
func (s *CustomInfo) ToXML() (string, error) {
	if s.info == "" {
		return "", makeError(ErrEmptyField, "Empty %s value", s.tag())
	}
	if s.infoType == "" {
		return "", makeError(ErrEmptyAttribute, "Empty attribute %s for custom-info", s.tag())
	}
	b := s.builder()
	b.Reset()
	b.openTagAttr(s.tag(), map[string]string{"info-type": s.infoType}, false)
	b.WriteString(s.info)
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *CustomInfo) tag() string {
	return "custom-info"
}
