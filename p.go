package gofbwriter

type p struct {
	text string
}

func (s *p) ToXML() (string, error) {
	return makeTag("p", sanitizeString(s.text)), nil
}

type empty struct {
}

func (s *empty) ToXML() (string, error) {
	return "<empty-line />", nil
}
