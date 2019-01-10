package gofbwriter

//A basic block of a book, can contain more child sections or textual content
type section struct {
	b          *builder
	title      *title      //Section's title
	epigraphs  []*epigraph //Epigraph(s) for this section
	image      *image      //Image to be displayed at the top of this section
	annotation *annotation //Annotation for this section, if any
	id         string
	sections   []*section //Child sections
	items      []fb
}

func (s *section) CreateTitle() *title {
	s.title = &title{}
	return s.title
}

func (s *section) CreateEpigraph() *epigraph {
	ep := &epigraph{}
	if s.epigraphs == nil {
		s.epigraphs = []*epigraph{ep}
	} else {
		s.epigraphs = append(s.epigraphs, ep)
	}
	return ep
}

func (s *section) CreateSectionImage() *image {
	s.image = &image{}
	return s.image
}

func (s *section) CreateAnnotation() *annotation {
	s.annotation = &annotation{}
	return s.annotation
}

func (s *section) CreateSection() (*section, error) {
	if s.items != nil && len(s.items) > 0 {
		return nil, makeError(ErrUnsupportedNestedItem, "Can't mix nested sections with formatted text in section")
	}
	item := &section{}
	if s.sections == nil {
		s.sections = []*section{item}
	} else {
		s.sections = append(s.sections, item)
	}
	return item, nil
}

func (s *section) AddSection(item *section) error {
	if s.items != nil && len(s.items) > 0 {
		return makeError(ErrUnsupportedNestedItem, "Can't mix nested sections with formatted text in section")
	}
	if s.sections == nil {
		s.sections = []*section{item}
	} else {
		s.sections = append(s.sections, item)
	}
	return nil
}

func (s *section) AddParagraph(text string) error {
	return s.AppendItem(&p{text: text})
}

func (s *section) CreateImage() (*image, error) {
	item := &image{}
	err := s.AppendItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *section) CreatePoem() (*poem, error) {
	item := &poem{}
	err := s.AppendItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (s *section) AddSubtitle(text string) error {
	return s.AppendItem(&p{text: text, tagName: "subtitle"})
}

func (s *section) CreateCite() (*cite, error) {
	item := &cite{}
	err := s.AppendItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *section) AddEmptyLine() error {
	return s.AppendItem(&empty{})
}

func (s *section) CreateTable() (*table, error) {
	item := &table{}
	err := s.AppendItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *section) AddImage(item *image) error {
	return s.AppendItem(item)
}

func (s *section) AddPoem(item *poem) error {
	return s.AppendItem(item)
}

func (s *section) AddCite(item *cite) error {
	return s.AppendItem(item)
}

func (s *section) AddTable(item *table) error {
	return s.AppendItem(item)
}

func (s *section) AppendItem(item fb) error { // nolint: gocyclo
	if s.sections != nil && len(s.sections) > 0 {
		return makeError(ErrUnsupportedNestedItem, "Can't mix nested sections with formatted text in section")
	}
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
	} else if _, ok := item.(*image); ok {
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
		s.items = []fb{item}
	} else {
		s.items = append(s.items, item)
	}
	return nil
}

func (s *section) ID() string {
	return s.id
}

func (s *section) SetID(id string) {
	s.id = id
}

func (s *section) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *section) ToXML() (string, error) {
	if s.IsEmpty() {
		return "", makeError(ErrEmptyField, "Section musts contain nested sections or formatted text")
	}
	b := s.builder()
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
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *section) makeItems() error {
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

func (s *section) makeSections() error {
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

func (s *section) makeAnnotation() error {
	if s.annotation != nil {
		str, err := s.annotation.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), s.annotation.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *section) makeImage() error {
	if s.image != nil {
		str, err := s.image.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), s.image.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *section) makeEpigraphs() error {
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

func (s *section) makeTitle() error {
	if s.title != nil {
		str, err := s.title.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), s.title.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *section) IsEmpty() bool {
	return (s.items == nil || len(s.items) == 0) && (s.sections == nil || len(s.sections) == 0)
}

func (s *section) tag() string {
	return "section"
}
