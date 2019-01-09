package gofbwriter

import (
	"encoding/base64"
)

//Any binary data that is required for the presentation of this book in base64 format. Currently only images are used.
type binary struct {
	b           *builder
	id          string
	contentType string
	data        []byte
}

func (s *binary) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
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
	b := s.builder()
	b.openTagAttr(s.tag(), map[string]string{"id": s.id, "content-type": s.contentType}, false)
	b.WriteString(base64.StdEncoding.EncodeToString(s.data))
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *binary) tag() string {
	return "binary"
}
