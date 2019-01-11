package gofbwriter

//Cite - a citation with an optional citation author at the end
type Cite struct {
	b           *builder
	id          string
	items       []Fb
	textAuthors []string
}

func (s *Cite) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//ID - get ID attribute
func (s *Cite) ID() string {
	return s.id
}

//SetID - set ID attribute
func (s *Cite) SetID(id string) {
	s.id = id
}

//TextAuthors - get list of authors
func (s *Cite) TextAuthors() []string {
	return s.textAuthors
}

//AddTextAuthor - add new text author to list of authors
func (s *Cite) AddTextAuthor(textAuthor string) {
	if s.textAuthors == nil {
		s.textAuthors = []string{textAuthor}
	} else {
		s.textAuthors = append(s.textAuthors, textAuthor)
	}
}

//AddParagraph - add new paragraph to cite
func (s *Cite) AddParagraph(par string) {
	item := &p{text: par}
	_ = s.AppendItem(item)
}

//CreatePoem - create and return new poem
func (s *Cite) CreatePoem() *Poem {
	item := &Poem{}
	_ = s.AppendItem(item)
	return item
}

//AddEmptyLine - add new empty line to cite
func (s *Cite) AddEmptyLine() {
	_ = s.AppendItem(&empty{})
}

//AddSubtitle - add new subtitle to cite
func (s *Cite) AddSubtitle(text string) {
	item := &p{tagName: "subtitle", text: text}
	_ = s.AppendItem(item)
}

//CreateTable - create and return new table
func (s *Cite) CreateTable() *Table {
	item := &Table{}
	_ = s.AppendItem(item)
	return item
}

//AppendItem - append existing item to cite
func (s *Cite) AppendItem(item Fb) error {
	pass := false
	if _, ok := item.(*p); ok {
		pass = true
	} else if _, ok := item.(*Poem); ok {
		pass = true
	} else if _, ok := item.(*p); ok {
		pass = true
	} else if _, ok := item.(*empty); ok {
		pass = true
	} else if _, ok := item.(*Table); ok {
		pass = true
	}
	if !pass {
		return makeError(ErrUnsupportedNestedItem, "Can't use type %T in cite", item)
	}
	if s.items == nil {
		s.items = []Fb{item}
	} else {
		s.items = append(s.items, item)
	}
	return nil
}

//ToXML - export to XML string
func (s *Cite) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
	b.openTagAttr(s.tag(), map[string]string{"id": s.id}, false)
	if s.items != nil {
		for _, item := range s.items {
			str, err := item.ToXML()
			if err != nil {
				return "", wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), item.tag())
			}
			b.WriteString(str)
		}
	}
	b.makeTags("text-author", s.textAuthors, true)
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *Cite) tag() string {
	return "cite"
}
