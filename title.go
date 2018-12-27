package gofbwriter

import (
	"fmt"
	"strings"
)

//A title, used in sections, poems and body elements
type title struct {
	items []fb
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
	var b strings.Builder
	fmt.Fprintf(&b, "<%s>\n", s.tag())
	for _, i := range s.items {
		str, err := i.ToXML()
		if err != nil {
			return "", wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), i.tag())
		}
		b.WriteString(str)
	}
	fmt.Fprintf(&b, "</%s>\n", s.tag())
	return b.String(), nil
}

func (s *title) tag() string {
	return "title"
}
