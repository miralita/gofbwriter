package gofbwriter

import (
	"fmt"
	"strings"
)

type customInfo struct {
	info     string
	infoType string
}

func (s *customInfo) ToXML() (string, error) {
	if s.info == "" {
		return "", makeError(ErrEmptyField, "Empty custom-info value")
	}
	if s.infoType == "" {
		return "", makeError(ErrEmptyAttribute, "Empty attribute info-type for custom-info")
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("<custom-info %s>\n", makeAttribute("info-type", s.infoType)))
	b.WriteString(s.info)
	b.WriteString("</custom-info>\n")
	return b.String(), nil
}
