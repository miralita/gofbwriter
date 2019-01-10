package gofbwriter

import (
	"time"
)

/*<xs:complexType name="poemType">
  <xs:annotation>
    <xs:documentation>A poem</xs:documentation>
  </xs:annotation>
  <xs:sequence>
    <xs:element name="title" type="titleType" minOccurs="0">
      <xs:annotation>
        <xs:documentation>Poem title</xs:documentation>
      </xs:annotation>
    </xs:element>
    <xs:element name="epigraph" type="epigraphType" minOccurs="0" maxOccurs="unbounded">
      <xs:annotation>
        <xs:documentation>Poem epigraph(s), if any</xs:documentation>
      </xs:annotation>
    </xs:element>
    <xs:choice maxOccurs="unbounded">
      <xs:element name="subtitle" type="pType"/>
      <xs:element name="stanza"></xs:element>
    </xs:choice>
    <xs:element name="text-author" type="pType" minOccurs="0" maxOccurs="unbounded"/>
    <xs:element name="date" type="dateType" minOccurs="0">
      <xs:annotation>
        <xs:documentation>Date this poem was written.</xs:documentation>
      </xs:annotation>
    </xs:element>
  </xs:sequence>
  <xs:attribute name="id" type="xs:ID" use="optional"/>
  <xs:attribute ref="xml:lang"/>
</xs:complexType>*/
//A poem
type poem struct {
	b *builder
	//Poem title
	title *title
	//Poem epigraph(s), if any
	epigraphs []*epigraph
	//Each poem should have at least one stanza. Stanzas are usually separated with empty lines by user agents.
	items   []fb //stanzas and subtitles
	authors []string
	date    *date //Date this poem was written.
	id      string
}

func (s *poem) ID() string {
	return s.id
}

func (s *poem) SetID(id string) {
	s.id = id
}

func (s *poem) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *poem) Date() *date {
	return s.date
}

func (s *poem) Authors() []string {
	return s.authors
}

func (s *poem) Items() []fb {
	return s.items
}

func (s *poem) Epigraphs() []*epigraph {
	return s.epigraphs
}

func (s *poem) Title() *title {
	return s.title
}

type subtitle string

func (s subtitle) builder() *builder {
	return &builder{}
}

func (s subtitle) ToXML() (string, error) {
	return s.builder().makeTag(s.tag(), sanitizeString(string(s))).String(), nil
}

func (s subtitle) tag() string {
	return "subtitle"
}

func (s *poem) SetDate(dt time.Time) {
	d := date(dt)
	s.date = &d
}

func (s *poem) CreateSubtitle(subtl string) {
	str := subtitle(subtl)
	s.addItem(str)
}

func (s *poem) CreateStanza() *stanza {
	st := &stanza{}
	s.addItem(st)
	return st
}

func (s *poem) CreateTitle() *title {
	t := &title{}
	s.addItem(t)
	return t
}

func (s *poem) AddAuthor(author string) {
	if s.authors == nil || len(s.authors) == 0 {
		s.authors = []string{author}
	} else {
		s.authors = append(s.authors, author)
	}
}

func (s *poem) CreateEpigraph() *epigraph {
	ep := &epigraph{}
	if s.epigraphs == nil || len(s.epigraphs) == 0 {
		s.epigraphs = []*epigraph{ep}
	} else {
		s.epigraphs = append(s.epigraphs, ep)
	}
	return ep
}

func (s *poem) ToXML() (string, error) {
	b := s.builder()
	b.openTagAttr(s.tag(), map[string]string{"id": s.id}, false)
	if err := s.makeTitle(); err != nil {
		return "", err
	}
	if err := s.makeEpigraphs(); err != nil {
		return "", err
	}
	if err := s.makeStanzas(); err != nil {
		return "", err
	}
	s.makeAuthor()
	if err := s.makeDate(); err != nil {
		return "", err
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *poem) addItem(item fb) {
	if s.items == nil || len(s.items) == 0 {
		s.items = []fb{item}
	} else {
		s.items = append(s.items, item)
	}
}

func (s *poem) makeEpigraphs() error {
	if s.epigraphs == nil || len(s.epigraphs) == 0 {
		return nil
	}
	for _, ep := range s.epigraphs {
		str, err := ep.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/epigraph", s.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *poem) makeTitle() error {
	if s.title == nil {
		return makeError(ErrEmptyField, "Empty required %s/title", s.tag())
	}
	str, err := s.title.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/title", s.tag())
	}
	s.builder().WriteString(str)
	return nil
}

func (s *poem) makeStanzas() error {
	if s.items == nil || len(s.items) == 0 {
		return makeError(ErrEmptyField, "Empty required %s/stanza", s.tag())
	}
	ok := false
	for _, i := range s.items {
		if !ok {
			_, ok = i.(*stanza)
		}
		str, err := i.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), i.tag())
		}
		s.builder().WriteString(str)
	}
	if !ok {
		return makeError(ErrEmptyField, "Each poem should have at least one stanza")
	}
	return nil
}

func (s *poem) makeAuthor() error {
	if s.authors == nil {
		return nil
	}
	s.builder().makeTags("text-author", s.authors, true)
	return nil
}

func (s *poem) makeDate() error {
	if s.date == nil {
		return nil
	}
	str, err := s.date.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), s.date.tag())
	}
	s.builder().WriteString(str)
	return nil
}

func (s *poem) tag() string {
	return "poem"
}
