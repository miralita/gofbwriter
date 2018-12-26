package gofbwriter

import (
	"strings"
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
type poem struct {
	title     *title
	epigraphs []*epigraph
	items     []fb //stanzas and subtitles
	authors   []string
	date      *date
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

func (s subtitle) ToXML() (string, error) {
	return makeTag(s.tag(), sanitizeString(string(s))), nil
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
	var b strings.Builder
	b.WriteString("<poem>\n")
	if err := s.makeTitle(&b); err != nil {
		return "", wrapError(err, ErrNestedEntity, "Can't make poem/title")
	}
	if err := s.makeEpigraphs(&b); err != nil {
		return "", wrapError(err, ErrNestedEntity, "Can't make poem/epigraph")
	}
	if err := s.makeStanzas(&b); err != nil {
		return "", wrapError(err, ErrNestedEntity, "Can't make poem/stanza")
	}
	if err := s.makeAuthor(&b); err != nil {
		return "", wrapError(err, ErrNestedEntity, "Can't make poem/text-author")
	}
	if err := s.makeDate(&b); err != nil {
		return "", wrapError(err, ErrNestedEntity, "Can't make poem/date")
	}
	b.WriteString("</poem>\n")
	return b.String(), nil
}

func (s *poem) addItem(item fb) {
	if s.items == nil || len(s.items) == 0 {
		s.items = []fb{item}
	} else {
		s.items = append(s.items, item)
	}
}

func (s *poem) makeEpigraphs(b *strings.Builder) error {
	if s.epigraphs == nil || len(s.epigraphs) == 0 {
		return nil
	}
	for _, ep := range s.epigraphs {
		str, err := ep.ToXML()
		if err != nil {
			return err
		}
		b.WriteString(str)
	}
	return nil
}

func (s *poem) makeTitle(b *strings.Builder) error {
	if s.title == nil {
		return makeError(ErrEmptyField, "Empty required poem/title")
	}
	str, err := s.title.ToXML()
	if err != nil {
		return err
	}
	b.WriteString(str)
	return nil
}

func (s *poem) makeStanzas(b *strings.Builder) error {
	if s.items == nil || len(s.items) == 0 {
		return makeError(ErrEmptyField, "Empty required poem/stanza")
	}
	for _, i := range s.items {
		str, err := i.ToXML()
		if err != nil {
			return err
		}
		b.WriteString(str)
	}
	return nil
}

func (s *poem) makeAuthor(b *strings.Builder) error {
	if s.authors == nil {
		return nil
	}
	b.WriteString(makeTagMulti("text-author", s.authors, true))
	return nil
}

func (s *poem) makeDate(b *strings.Builder) error {
	if s.date == nil {
		return nil
	}
	str, err := s.date.ToXML()
	if err != nil {
		return err
	}
	b.WriteString(str)
	return nil
}

func (s *poem) tag() string {
	return "poem"
}
