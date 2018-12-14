package gofbwriter

//A basic paragraph, may include simple formatting inside
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
