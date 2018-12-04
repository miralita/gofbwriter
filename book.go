package gofbwriter

type book struct {
	stylesheet  *stylesheet
	description *description
	body        *body
	binary      []binary
}

func (s *book) SetStylesheet(ctype, data string) (*stylesheet, error) {
	if ctype == "" {
		ctype = "text/css"
	}
	s.stylesheet = &stylesheet{ctype, data, s}
	return s.stylesheet, nil
}

func (s *book) Stylesheet() *stylesheet {
	if s.stylesheet == nil {
		s.stylesheet = &stylesheet{book: s}
	}
	return s.stylesheet
}

func (s *book) Description() *description {
	if s.description == nil {
		s.description = &description{book: s}
	}
	return s.description
}
