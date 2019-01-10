package gofbwriter

//A cut-down version of &lt;section&gt; used in annotations
type annotation struct {
	b       *builder
	tagName string
	items   []fb
	id      string
}

func (s *annotation) ID() string {
	return s.id
}

func (s *annotation) SetID(id string) {
	s.id = id
}

func (s *annotation) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *annotation) ToXML() (string, error) {
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
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *annotation) AddParagraph(text string) {
	par := &p{text: text}
	s.AppendItem(par)
}

func (s *annotation) CreatePoem() *poem {
	item := &poem{}
	s.AppendItem(item)
	return item
}

func (s *annotation) CreateCite() *cite {
	item := &cite{}
	s.AppendItem(item)
	return item
}

func (s *annotation) AddSubtitle(text string) {
	item := &p{tagName: "subtitle", text: text}
	s.AppendItem(item)
}

func (s *annotation) CreateTable() *table {
	item := &table{}
	s.AppendItem(item)
	return item
}

func (s *annotation) AddEmptyLine() {
	s.AppendItem(&empty{})
}

func (s *annotation) AppendItem(item fb) error {
	pass := false
	if _, ok := item.(*p); ok {
		pass = true
	} else if _, ok := item.(*poem); ok {
		pass = true
	} else if _, ok := item.(*cite); ok {
		pass = true
	} else if _, ok := item.(*empty); ok {
		pass = true
	} else if _, ok := item.(*table); ok {
		pass = true
	}
	if !pass {
		return makeError(ErrUnsupportedNestedItem, "Can't use type %T in annotation", item)
	}
	if s.items == nil {
		s.items = []fb{item}
	} else {
		s.items = append(s.items, item)
	}
	return nil
}

func (s *annotation) tag() string {
	if s.tagName == "" {
		s.tagName = "annotation"
	}
	return s.tagName
}
