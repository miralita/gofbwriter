package gofbwriter

//A basic paragraph, may include simple formatting inside
type p struct {
	tagName string
	text    string
}

func (s *p) ToXML() (string, error) {
	if s.tagName == "" {
		s.tagName = "p"
	}
	return makeTag(s.tagName, sanitizeString(s.text)), nil
}

func (s *p) tag() string {
	return "p"
}

type empty struct {
}

func (s *empty) ToXML() (string, error) {
	return "<empty-line />", nil
}

func (s *empty) tag() string {
	return "empty-line"
}
