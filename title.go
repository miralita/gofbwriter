package gofbwriter

import "strings"

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
		return "<title />\n", nil
	}
	var b strings.Builder
	b.WriteString("<title>")
	for _, i := range s.items {
		str, err := i.ToXML()
		if err != nil {
			return "", wrapError(err, ErrNestedEntity, "Can't make title nested elements")
		}
		b.WriteString(str)
	}
	b.WriteString("</title>\n")
	return b.String(), nil
}
