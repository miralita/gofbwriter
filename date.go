package gofbwriter

import (
	"fmt"
	"time"
)

//Date - A human readable date, maybe not exact, with an optional computer readable variant
type Date time.Time

func (s *Date) builder() *builder {
	return nil
}

//ToXML - export to XML string
func (s *Date) ToXML() (string, error) {
	dt := time.Time(*s).String()
	return fmt.Sprintf(`<%s value="%s">%s</%s>`, s.tag(), dt, dt, s.tag()), nil
}

func (s *Date) tag() string {
	return "date"
}
