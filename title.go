package gofbwriter

import (
	"fmt"
)

//A title, used in sections, poems and body elements
type title struct {
	b     *builder
	items []fb
}

func (s *title) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *title) CreateParagraph(str string) {
	s.appendItem(&p{text: str})
}

func (s *title) appendItem(i fb) {
	if s.items == nil {
		s.items = []fb{i}
	} else {
		s.items = append(s.items, i)
	}
}

func (s *title) CreateEmptyline() {
	s.appendItem(&empty{})
}

func (s *title) ToXML() (string, error) {
	if s.items == nil || len(s.items) == 0 {
		return fmt.Sprintf("<%s />\n", s.tag()), nil
	}
	b := s.builder()
	b.Reset()
	b.openTag(s.tag())
	for _, i := range s.items {
		str, err := i.ToXML()
		if err != nil {
			return "", wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), i.tag())
		}
		b.WriteString(str)
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *title) tag() string {
	return "title"
}
