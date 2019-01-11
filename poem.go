package gofbwriter

import (
	"time"
)

//Poem - a poem
type Poem struct {
	b *builder
	//Poem title
	title *Title
	//Poem epigraph(s), if any
	epigraphs []*Epigraph
	//Each poem should have at least one stanza. Stanzas are usually separated with empty lines by user agents.
	items   []Fb //stanzas and subtitles
	authors []string
	date    *Date //Date this poem was written.
	id      string
}

//ID - get id attribute
func (s *Poem) ID() string {
	return s.id
}

//SetID - set id attribute
func (s *Poem) SetID(id string) {
	s.id = id
}

func (s *Poem) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Date - get date this poem was written
func (s *Poem) Date() *Date {
	return s.date
}

//Authors - get authors
func (s *Poem) Authors() []string {
	return s.authors
}

//Items - get list of nested items
func (s *Poem) Items() []Fb {
	return s.items
}

//Epigraphs - get list of epigraphs
func (s *Poem) Epigraphs() []*Epigraph {
	return s.epigraphs
}

//Title - get Title
func (s *Poem) Title() *Title {
	return s.title
}

type subtitle string

func (s subtitle) builder() *builder {
	return &builder{}
}

//ToXML - export to XML string
func (s subtitle) ToXML() (string, error) {
	return s.builder().makeTag(s.tag(), sanitizeString(string(s))).String(), nil
}

func (s subtitle) tag() string {
	return "subtitle"
}

//SetDate - set date from time.Time. Old Date will be dropped
func (s *Poem) SetDate(dt time.Time) {
	d := Date(dt)
	s.date = &d
}

//AddSubtitle - add new sibtitle to list of items
func (s *Poem) AddSubtitle(subtl string) {
	str := subtitle(subtl)
	s.addItem(str)
}

//CreateStanza - create new Stanza, add it to list of items and return
func (s *Poem) CreateStanza() *Stanza {
	st := &Stanza{}
	s.addItem(st)
	return st
}

//CreateTitle - create new Title, add it to list of items and return
func (s *Poem) CreateTitle() *Title {
	t := &Title{}
	s.addItem(t)
	return t
}

//AddAuthor - add author to list of authors
func (s *Poem) AddAuthor(author string) {
	if s.authors == nil || len(s.authors) == 0 {
		s.authors = []string{author}
	} else {
		s.authors = append(s.authors, author)
	}
}

//CreateEpigraph - create new Epigraph, add it to list of epigraphs and return
func (s *Poem) CreateEpigraph() *Epigraph {
	ep := &Epigraph{}
	if s.epigraphs == nil || len(s.epigraphs) == 0 {
		s.epigraphs = []*Epigraph{ep}
	} else {
		s.epigraphs = append(s.epigraphs, ep)
	}
	return ep
}

//ToXML - export to XML string
func (s *Poem) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
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

func (s *Poem) addItem(item Fb) {
	if s.items == nil || len(s.items) == 0 {
		s.items = []Fb{item}
	} else {
		s.items = append(s.items, item)
	}
}

func (s *Poem) makeEpigraphs() error {
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

func (s *Poem) makeTitle() error {
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

func (s *Poem) makeStanzas() error {
	if s.items == nil || len(s.items) == 0 {
		return makeError(ErrEmptyField, "Empty required %s/stanza", s.tag())
	}
	ok := false
	for _, i := range s.items {
		if !ok {
			_, ok = i.(*Stanza)
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

func (s *Poem) makeAuthor() error {
	if s.authors == nil {
		return nil
	}
	s.builder().makeTags("text-author", s.authors, true)
	return nil
}

func (s *Poem) makeDate() error {
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

func (s *Poem) tag() string {
	return "poem"
}
