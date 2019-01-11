package gofbwriter

import (
	"fmt"
)

//Title - a title, used in sections, poems and body elements
type Title struct {
	b     *builder
	items []Fb
}

func (s *Title) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//AddParagraph - add new paragraph
func (s *Title) AddParagraph(str string) {
	s.appendItem(&p{text: str})
}

func (s *Title) appendItem(i Fb) {
	if s.items == nil {
		s.items = []Fb{i}
	} else {
		s.items = append(s.items, i)
	}
}

//AddEmptyLine - add new empty line
func (s *Title) AddEmptyLine() {
	s.appendItem(&empty{})
}

//ToXML - export to XML string
func (s *Title) ToXML() (string, error) {
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

func (s *Title) tag() string {
	return "title"
}
