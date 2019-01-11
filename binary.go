package gofbwriter

import (
	"encoding/base64"
)

//Binary - any binary data that is required for the presentation of this book in base64 format. Currently only images are used.
type Binary struct {
	b           *builder
	id          string
	contentType string
	data        []byte
}

func (s *Binary) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Data - get binary data
func (s *Binary) Data() []byte {
	return s.data
}

//ContentType - get content-type attribute
func (s *Binary) ContentType() string {
	return s.contentType
}

//ID - get ID attribute
func (s *Binary) ID() string {
	return s.id
}

//ToXML - export to XML string
func (s *Binary) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
	b.openTagAttr(s.tag(), map[string]string{"id": s.id, "content-type": s.contentType}, false)
	b.WriteString(base64.StdEncoding.EncodeToString(s.data))
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *Binary) tag() string {
	return "binary"
}
