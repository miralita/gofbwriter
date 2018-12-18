package gofbwriter

type book struct {
	//This element contains an arbitrary stylesheet that is intepreted by a some processing programs, e.g. text/css stylesheets can be used by XSLT stylesheets to generate better looking html
	stylesheets []*stylesheet
	description *description
	body        *body
	binary      []binary
}

func (s *book) CreateStylesheet(ctype, data string) (*stylesheet, error) {
	if ctype == "" {
		ctype = "text/css"
	}
	st := &stylesheet{ctype, data, s}
	if s.stylesheets == nil {
		s.stylesheets = []*stylesheet{st}
	} else {
		s.stylesheets = append(s.stylesheets, st)
	}
	return st, nil
}

func (s *book) Description() *description {
	if s.description == nil {
		s.description = &description{book: s}
	}
	return s.description
}
