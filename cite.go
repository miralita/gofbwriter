package gofbwriter

//A citation with an optional citation author at the end
type cite struct {
	b           *builder
	id          string
	items       []fb
	textAuthors []string
}

func (s *cite) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *cite) ID() string {
	return s.id
}

func (s *cite) SetID(id string) {
	s.id = id
}

func (s *cite) TextAuthors() []string {
	return s.textAuthors
}

func (s *cite) AddTextAuthor(textAuthor string) {
	if s.textAuthors == nil {
		s.textAuthors = []string{textAuthor}
	} else {
		s.textAuthors = append(s.textAuthors, textAuthor)
	}
}

func (s *cite) AddParagraph(par string) *p {
	item := &p{text: par}
	_ = s.AppendItem(item)
	return item
}

func (s *cite) CreatePoem() *poem {
	item := &poem{}
	_ = s.AppendItem(item)
	return item
}

func (s *cite) CreateEmptyLine() {
	_ = s.AppendItem(&empty{})
}

func (s *cite) CreateSubtitle(text string) *p {
	item := &p{tagName: "subtitle", text: text}
	_ = s.AppendItem(item)
	return item
}

func (s *cite) CreateTable() *table {
	item := &table{}
	_ = s.AppendItem(item)
	return item
}

func (s *cite) AppendItem(item fb) error {
	pass := false
	if _, ok := item.(*p); ok {
		pass = true
	} else if _, ok := item.(*poem); ok {
		pass = true
	} else if _, ok := item.(*p); ok {
		pass = true
	} else if _, ok := item.(*empty); ok {
		pass = true
	} else if _, ok := item.(*table); ok {
		pass = true
	}
	if !pass {
		return makeError(ErrUnsupportedNestedItem, "Can't use type %T in cite", item)
	}
	if s.items == nil {
		s.items = []fb{item}
	} else {
		s.items = append(s.items, item)
	}
	return nil
}

func (s *cite) ToXML() (string, error) {
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

func (s *cite) tag() string {
	return "cite"
}
