package gofbwriter

import (
	"fmt"
	"strings"
)

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
	var b strings.Builder
	b.WriteString(fmt.Sprintf("<%s %s>\n", s.tag(), makeAttribute("info-type", s.infoType)))
	b.WriteString(s.info)
	fmt.Fprintf(&b, "</%s>\n", s.tag())
	return b.String(), nil
}

func (s *customInfo) tag() string {
	return "custom-info"
}
