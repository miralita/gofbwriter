package gofbwriter

import (
	"fmt"
)

//Epigraph - An epigraph
type Epigraph struct {
	b           *builder
	textAuthors []string
	items       []Fb
	id          string
}

//ID - get ID attribute
func (s *Epigraph) ID() string {
	return s.id
}

//SetID - set ID attribute
func (s *Epigraph) SetID(id string) {
	s.id = id
}

func (s *Epigraph) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//TextAuthors - get list of authors
func (s *Epigraph) TextAuthors() []string {
	return s.textAuthors
}

//AddAuthor - add author name to the list of authors
func (s *Epigraph) AddAuthor(author string) {
	if s.textAuthors == nil {
		s.textAuthors = []string{author}
	} else {
		s.textAuthors = append(s.textAuthors, author)
	}
}

//AddParagraph - add new paragraph to epigraph
func (s *Epigraph) AddParagraph(str string) {
	p := &p{text: str}
	_ = s.AppendItem(p)
}

//AddEmptyLine - add empty-line tag
func (s *Epigraph) AddEmptyLine() {
	_ = s.AppendItem(&empty{})
}

//CreatePoem - create new Poem, add it to list and return
func (s *Epigraph) CreatePoem() *Poem {
	p := &Poem{}
	_ = s.AppendItem(p)
	return p
}

//CreateCite - create new Cite, add it to list and return
func (s *Epigraph) CreateCite() *Cite {
	p := &Cite{}
	_ = s.AppendItem(p)
	return p
}

//AppendItem - append item to list of items
func (s *Epigraph) AppendItem(item Fb) error {
	pass := false
	if _, ok := item.(*p); ok {
		pass = true
	} else if _, ok := item.(*Poem); ok {
		pass = true
	} else if _, ok := item.(*Cite); ok {
		pass = true
	} else if _, ok := item.(*empty); ok {
		pass = true
	}
	if !pass {
		return makeError(ErrUnsupportedNestedItem, "Can't use type %T in epigraph", item)
	}
	if s.items == nil {
		s.items = []Fb{item}
	} else {
		s.items = append(s.items, item)
	}
	return nil
}

//ToXML - export to XML string
func (s *Epigraph) ToXML() (string, error) {
	if (s.items == nil || len(s.items) == 0) && (s.textAuthors == nil || len(s.textAuthors) == 0) {
		return fmt.Sprintf("<%s />\n", s.tag()), nil
	}
	b := s.builder()
	b.Reset()
	b.openTagAttr(s.tag(), map[string]string{"id": s.id}, false)
	if s.items != nil {
		for _, i := range s.items {
			str, err := i.ToXML()
			if err != nil {
				return "", wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), i.tag())
			}
			b.WriteString(str)
		}
	}
	if s.textAuthors != nil {
		for _, i := range s.textAuthors {
			s.builder().makeTag("text-author", sanitizeString(i))
		}
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *Epigraph) tag() string {
	return "epigraph"
}
