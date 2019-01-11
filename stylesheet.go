package gofbwriter

//StyleSheet - This element contains an arbitrary stylesheet that is intepreted by a some processing programs, e.g. text/css stylesheets can be used by XSLT stylesheets to generate better looking html
type StyleSheet struct {
	b     *builder
	ctype string
	data  string
}

func (s *StyleSheet) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Set - set content-type and data for this stylesheet
func (s *StyleSheet) Set(ctype, data string) {
	s.ctype = ctype
	s.data = data
}

//ToXML - export to XML string
func (s *StyleSheet) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
	b.openTagAttr(s.tag(), map[string]string{"type": s.ctype}, false)
	b.WriteString(s.data)
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *StyleSheet) tag() string {
	return "stylesheet"
}
