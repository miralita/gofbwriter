package gofbwriter

//Section - a basic block of a book, can contain more child sections or textual content
type Section struct {
	b          *builder
	title      *Title      //Section's title
	epigraphs  []*Epigraph //Epigraph(s) for this section
	image      *Image      //Image to be displayed at the top of this section
	annotation *Annotation //Annotation for this section, if any
	id         string
	sections   []*Section //Child sections
	items      []Fb
}

//CreateTitle - create and return new Title. Old title will be dropped
func (s *Section) CreateTitle() *Title {
	s.title = &Title{}
	return s.title
}

//CreateEpigraph - create new Epigraph, add it to list and return
func (s *Section) CreateEpigraph() *Epigraph {
	ep := &Epigraph{}
	if s.epigraphs == nil {
		s.epigraphs = []*Epigraph{ep}
	} else {
		s.epigraphs = append(s.epigraphs, ep)
	}
	return ep
}

//CreateSectionImage - create and return new main image of the section. Old Image will be dropped
func (s *Section) CreateSectionImage() *Image {
	s.image = &Image{}
	return s.image
}

//CreateAnnotation - create and return new Annotation. Old annotation will be dropped
func (s *Section) CreateAnnotation() *Annotation {
	s.annotation = &Annotation{}
	return s.annotation
}

//CreateSection - create  new Section, add it to list and return.
func (s *Section) CreateSection() (*Section, error) {
	if s.items != nil && len(s.items) > 0 {
		return nil, makeError(ErrUnsupportedNestedItem, "Can't mix nested sections with formatted text in section")
	}
	item := &Section{}
	if s.sections == nil {
		s.sections = []*Section{item}
	} else {
		s.sections = append(s.sections, item)
	}
	return item, nil
}

//AddSection - add existing section to list of sections
func (s *Section) AddSection(item *Section) error {
	if s.items != nil && len(s.items) > 0 {
		return makeError(ErrUnsupportedNestedItem, "Can't mix nested sections with formatted text in section")
	}
	if s.sections == nil {
		s.sections = []*Section{item}
	} else {
		s.sections = append(s.sections, item)
	}
	return nil
}

//AddParagraph - add paragraph to formatted text block
func (s *Section) AddParagraph(text string) error {
	return s.AppendItem(&p{text: text})
}

//CreateImage - create new Image, add it to formatted text block and return
func (s *Section) CreateImage() (*Image, error) {
	item := &Image{}
	err := s.AppendItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//CreatePoem - create new Poem, add it to formatted text block and return
func (s *Section) CreatePoem() (*Poem, error) {
	item := &Poem{}
	err := s.AppendItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil

}

//AddSubtitle - add subtitle to formatted text block
func (s *Section) AddSubtitle(text string) error {
	return s.AppendItem(&p{text: text, tagName: "subtitle"})
}

//CreateCite - create new Cite, add it to formatted text block and  return
func (s *Section) CreateCite() (*Cite, error) {
	item := &Cite{}
	err := s.AppendItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//AddEmptyLine - add empty-line to formatted text block
func (s *Section) AddEmptyLine() error {
	return s.AppendItem(&empty{})
}

//CreateTable - create new Table, add it to formatted text block and return
func (s *Section) CreateTable() (*Table, error) {
	item := &Table{}
	err := s.AppendItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//AddImage - add existing Image to formatted text block
func (s *Section) AddImage(item *Image) error {
	return s.AppendItem(item)
}

//AddPoem - add existing Poem to formatted text block
func (s *Section) AddPoem(item *Poem) error {
	return s.AppendItem(item)
}

//AddCite - add existing Cite to formatted text block
func (s *Section) AddCite(item *Cite) error {
	return s.AppendItem(item)
}

//AddTable - add existing Table to formatted text block
func (s *Section) AddTable(item *Table) error {
	return s.AppendItem(item)
}

//AppendItem - add existing element to formatted text block
func (s *Section) AppendItem(item Fb) error { // nolint: gocyclo
	if s.sections != nil && len(s.sections) > 0 {
		return makeError(ErrUnsupportedNestedItem, "Can't mix nested sections with formatted text in section")
	}
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
	} else if _, ok := item.(*Image); ok {
		if s.items != nil && len(s.items) > 0 {
			pass = true
		} else {
			return makeError(ErrUnsupportedNestedItem, "Can't use type %T as first element in section", item)
		}
	}
	if !pass {
		return makeError(ErrUnsupportedNestedItem, "Can't use type %T in section", item)
	}
	if s.items == nil {
		s.items = []Fb{item}
	} else {
		s.items = append(s.items, item)
	}
	return nil
}

//ID - get ID attribute
func (s *Section) ID() string {
	return s.id
}

//SetID - set ID attribute
func (s *Section) SetID(id string) {
	s.id = id
}

func (s *Section) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//ToXML - export to XML string
func (s *Section) ToXML() (string, error) {
	if s.IsEmpty() {
		return "", makeError(ErrEmptyField, "Section musts contain nested sections or formatted text")
	}
	b := s.builder()
	b.Reset()
	b.openTagAttr(s.tag(), map[string]string{"id": s.id}, false)
	if err := s.makeTitle(); err != nil {
		return "", err
	}
	if err := s.makeEpigraphs(); err != nil {
		return "", err
	}
	if err := s.makeImage(); err != nil {
		return "", err
	}
	if err := s.makeAnnotation(); err != nil {
		return "", err
	}
	if err := s.makeSections(); err != nil {
		return "", err
	}
	if err := s.makeItems(); err != nil {
		return "", err
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *Section) makeItems() error {
	if s.items != nil {
		for _, item := range s.items {
			str, err := item.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), item.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *Section) makeSections() error {
	if s.sections != nil {
		for _, item := range s.sections {
			str, err := item.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), item.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *Section) makeAnnotation() error {
	if s.annotation != nil {
		str, err := s.annotation.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), s.annotation.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *Section) makeImage() error {
	if s.image != nil {
		str, err := s.image.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), s.image.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *Section) makeEpigraphs() error {
	if s.epigraphs != nil {
		for _, item := range s.epigraphs {
			str, err := item.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), item.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *Section) makeTitle() error {
	if s.title != nil {
		str, err := s.title.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), s.title.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

//IsEmpty - check if section doesn't contain any nested sections or formatted text
func (s *Section) IsEmpty() bool {
	return (s.items == nil || len(s.items) == 0) && (s.sections == nil || len(s.sections) == 0)
}

func (s *Section) tag() string {
	return "section"
}
