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
	b.WriteString("<stylesheet")
	if s.ctype != "" {
		fmt.Fprintf(&b, ` type="%s"`, s.ctype)
	}
	b.WriteString(">\n")
	b.WriteString(s.data)
	b.WriteString("\n</stylesheet>")
	return b.String(), nil
}

func (s *stylesheet) tag() string {
	return "stylesheet"
}
