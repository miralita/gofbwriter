package gofbwriter

import (
	"encoding/base64"
	"strings"
)

//Any binary data that is required for the presentation of this book in base64 format. Currently only images are used.
type binary struct {
	id          string
	contentType string
	data        []byte
}

func (s *binary) Data() []byte {
	return s.data
}

func (s *binary) ContentType() string {
	return s.contentType
}

func (s *binary) ID() string {
	return s.id
}

func (s *binary) ToXML() (string, error) {
	var b strings.Builder
	b.WriteString("<binary ")
	b.WriteString(makeAttribute("id", s.id))
	b.WriteString(" ")
	b.WriteString(makeAttribute("content-type", s.contentType))
	b.WriteString(">")
	b.WriteString(base64.StdEncoding.EncodeToString(s.data))
	b.WriteString("</binary>\n")
	return b.String(), nil
}

func (s *binary) tag() string {
	return "binary"
}
