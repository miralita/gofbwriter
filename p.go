package gofbwriter

import (
	"fmt"
)

//A basic paragraph, may include simple formatting inside
type p struct {
	b       *builder
	tagName string
	text    string
}

func (s *p) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *p) ToXML() (string, error) {
	s.builder().makeTag(s.tag(), sanitizeString(s.text))
	return s.b.String(), nil
}

func (s *p) tag() string {
	return "p"
}

type empty struct {
	b *builder
}

func (s *empty) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *empty) ToXML() (string, error) {
	return fmt.Sprintf("<%s />\n", s.tag()), nil
}

func (s *empty) tag() string {
	return "empty-line"
}
