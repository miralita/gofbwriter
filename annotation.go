package gofbwriter

//Annotation - A cut-down version of &lt;section&gt; used in annotations
type Annotation struct {
	b       *builder
	tagName string
	items   []Fb
	id      string
}

//ID - get id attribute
func (s *Annotation) ID() string {
	return s.id
}

//SetID - set id attribute
func (s *Annotation) SetID(id string) {
	s.id = id
}

func (s *Annotation) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//ToXML - export to XML string
func (s *Annotation) ToXML() (string, error) {
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

//AddParagraph - add new paragraph
func (s *Annotation) AddParagraph(text string) {
	par := &p{text: text}
	s.AppendItem(par)
}

//CreatePoem - create new Poem, add it to list of items and return
func (s *Annotation) CreatePoem() *Poem {
	item := &Poem{}
	s.AppendItem(item)
	return item
}

//CreateCite - create new Cite, add it to list of items and return
func (s *Annotation) CreateCite() *Cite {
	item := &Cite{}
	s.AppendItem(item)
	return item
}

//AddSubtitle - add new subtitle to list of items
func (s *Annotation) AddSubtitle(text string) {
	item := &p{tagName: "subtitle", text: text}
	s.AppendItem(item)
}

//CreateTable - create new Table, add it to list of items and return
func (s *Annotation) CreateTable() *Table {
	item := &Table{}
	s.AppendItem(item)
	return item
}

//AddEmptyLine - add new empty line to list of items
func (s *Annotation) AddEmptyLine() {
	s.AppendItem(&empty{})
}

//AppendItem - append existing item to list of items
func (s *Annotation) AppendItem(item Fb) error {
	pass := false
	if _, ok := item.(*p); ok {
		pass = true
	} else if _, ok := item.(*Poem); ok {
		pass = true
	} else if _, ok := item.(*Cite); ok {
		pass = true
	} else if _, ok := item.(*empty); ok {
		pass = true
	} else if _, ok := item.(*Table); ok {
		pass = true
	}
	if !pass {
		return makeError(ErrUnsupportedNestedItem, "Can't use type %T in annotation", item)
	}
	if s.items == nil {
		s.items = []Fb{item}
	} else {
		s.items = append(s.items, item)
	}
	return nil
}

func (s *Annotation) tag() string {
	if s.tagName == "" {
		s.tagName = "annotation"
	}
	return s.tagName
}
