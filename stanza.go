package gofbwriter

//Stanza - Each poem should have at least one stanza. Stanzas are usually separated with empty lines by user agents.
type Stanza struct {
	b        *builder
	title    *Title
	subtitle string
	lines    []string //An individual line in a stanza
}

func (s *Stanza) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Lines - get lines of stanza
func (s *Stanza) Lines() []string {
	return s.lines
}

//Subtitle - get subtitile
func (s *Stanza) Subtitle() string {
	return s.subtitle
}

//Title - get title
func (s *Stanza) Title() *Title {
	return s.title
}

//CreateTitle - create and return new Title; old value will be dropped
func (s *Stanza) CreateTitle() *Title {
	t := &Title{}
	s.title = t
	return t
}

//SetSubtitle - set subtitle
func (s *Stanza) SetSubtitle(subtl string) {
	s.subtitle = subtl
}

//AddLine - add new line to stanza
func (s *Stanza) AddLine(line string) {
	if s.lines == nil {
		s.lines = []string{line}
	} else {
		s.lines = append(s.lines, line)
	}
}

//ToXML - export to XML string
func (s *Stanza) ToXML() (string, error) {
	if s.lines == nil || len(s.lines) == 0 {
		return "", makeError(ErrEmptyField, "Empty required %s/v", s.tag())
	}
	b := s.builder()
	b.Reset()
	b.openTag(s.tag())
	if s.title != nil {
		str, err := s.ToXML()
		if err != nil {
			return "", wrapError(err, ErrNestedEntity, "Can't make %s/title", s.tag())
		}
		b.WriteString(str)
	}
	b.makeTag("subtitle", sanitizeString(s.subtitle))
	b.makeTags("v", s.lines, true)
	b.closeTag(s.tag())

	return b.String(), nil
}

func (s *Stanza) tag() string {
	return "stanza"
}
