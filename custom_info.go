package gofbwriter

type customInfo struct {
	b        *builder
	info     string
	infoType string
}

func (s *customInfo) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *customInfo) ToXML() (string, error) {
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

func (s *customInfo) tag() string {
	return "custom-info"
}
