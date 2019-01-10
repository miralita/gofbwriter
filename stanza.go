package gofbwriter

/*<xs:element name="stanza">
  <xs:annotation>
    <xs:documentation>Each poem should have at least one stanza. Stanzas are usually separated with empty lines by user agents.</xs:documentation>
  </xs:annotation>
  <xs:complexType>
    <xs:sequence>
      <xs:element name="title" type="titleType" minOccurs="0"/>
      <xs:element name="subtitle" type="pType" minOccurs="0"/>
      <xs:element name="v" type="pType" maxOccurs="unbounded">
        <xs:annotation>
          <xs:documentation>An individual line in a stanza</xs:documentation>
        </xs:annotation>
      </xs:element>
    </xs:sequence>
    <xs:attribute ref="xml:lang"/>
  </xs:complexType>
</xs:element>*/
type stanza struct {
	b        *builder
	title    *title
	subtitle string
	lines    []string //An individual line in a stanza
}

func (s *stanza) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *stanza) Lines() []string {
	return s.lines
}

func (s *stanza) Subtitle() string {
	return s.subtitle
}

func (s *stanza) Title() *title {
	return s.title
}

func (s *stanza) CreateTitle() *title {
	t := &title{}
	s.title = t
	return t
}

func (s *stanza) SetSubtitle(subtl string) {
	s.subtitle = subtl
}

func (s *stanza) AddLine(line string) {
	if s.lines == nil {
		s.lines = []string{line}
	} else {
		s.lines = append(s.lines, line)
	}
}

func (s *stanza) ToXML() (string, error) {
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

func (s *stanza) tag() string {
	return "stanza"
}
