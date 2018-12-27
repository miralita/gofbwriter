package gofbwriter

import (
	"fmt"
	"strings"
)

type stylesheet struct {
	ctype string
	data  string
	book  *book
}

func (s *stylesheet) Set(ctype, data string) {
	s.ctype = ctype
	s.data = data
}

func (s *stylesheet) ToXML() (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "<%s", s.tag())
	if s.ctype != "" {
		fmt.Fprintf(&b, ` type="%s"`, s.ctype)
	}
	fmt.Fprintf(&b, ">\n%s\n</%s>\n", s.data, s.tag())
	return b.String(), nil
}

func (s *stylesheet) tag() string {
	return "stylesheet"
}
