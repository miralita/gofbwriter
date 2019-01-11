package gofbwriter

//Body - main content of the book, multiple bodies are used for additional information, like footnotes, that do not appear in the main book flow (extended from this class). The first body is presented to the reader by default, and content in the other bodies should be accessible by hyperlinks.
type Body struct {
	b         *builder
	image     *Image      //Image to be displayed at the top of this section
	title     *Title      //A fancy title for the entire book, should be used if the simple text version in &lt;description&gt; is not adequate, e.g. the book title has multiple paragraphs and/or character styles
	epigraphs []*Epigraph //Epigraph(s) for the entire book, if any
	sections  []*Section
	name      string
}

func (s *Body) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Sections - get list of sections
func (s *Body) Sections() []*Section {
	return s.sections
}

//AddSection - add existing section to list of sections
func (s *Body) AddSection(sec *Section) {
	if s.sections == nil {
		s.sections = []*Section{sec}
	} else {
		s.sections = append(s.sections, sec)
	}
}

//CreateSection - create new Section and add it to list of sections, return created struct
func (s *Body) CreateSection() *Section {
	sec := &Section{}
	s.AddSection(sec)
	return sec
}

//Epigraphs - get list of epigraphs
func (s *Body) Epigraphs() []*Epigraph {
	return s.epigraphs
}

//AddEpigraph - add existing Epigraph to list
func (s *Body) AddEpigraph(ep *Epigraph) {
	if s.epigraphs == nil {
		s.epigraphs = []*Epigraph{ep}
	} else {
		s.epigraphs = append(s.epigraphs, ep)
	}
}

//CreateEpigraph - create new Epigraph, add it to list and return
func (s *Body) CreateEpigraph() *Epigraph {
	ep := &Epigraph{}
	s.AddEpigraph(ep)
	return ep
}

//Title - get Title
func (s *Body) Title() *Title {
	return s.title
}

//SetTitle - set existing Title
func (s *Body) SetTitle(title *Title) {
	s.title = title
}

//CreateTitle - create and return new Title. Old title will be dropped
func (s *Body) CreateTitle() *Title {
	s.title = &Title{}
	return s.title
}

//Image - get Image
func (s *Body) Image() *Image {
	return s.image
}

//CreateImage - create and return new Image. Old image will be dropped
func (s *Body) CreateImage() *Image {
	img := &Image{}
	s.image = img
	return img
}

//SetImage - set existing Image
func (s *Body) SetImage(image *Image) {
	s.image = image
}

//ToXML - export to XML string
func (s *Body) ToXML() (string, error) {
	if s.sections == nil || len(s.sections) == 0 {
		return "", makeError(ErrEmptyField, "Empty required field %s/section", s.tag())
	}
	b := s.builder()
	b.Reset()
	b.openTagAttr(s.tag(), map[string]string{"name": s.name}, false)
	if err := s.serializeImage(); err != nil {
		return "", err
	}
	if err := s.serializeTitle(); err != nil {
		return "", err
	}
	if err := s.serializeEpigraphs(); err != nil {
		return "", err
	}
	if err := s.serializeSections(); err != nil {
		return "", err
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *Body) serializeImage() error {
	if s.image != nil {
		i, err := s.image.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/image", s.tag())
		}
		s.builder().WriteString(i)
	}
	return nil
}

func (s *Body) serializeTitle() error {
	if s.title != nil {
		t, err := s.title.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/title", s.tag())
		}
		s.builder().WriteString(t)
	}
	return nil
}

func (s *Body) serializeEpigraphs() error {
	if s.epigraphs != nil && len(s.epigraphs) > 0 {
		for _, ep := range s.epigraphs {
			str, err := ep.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/epigraph", s.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *Body) serializeSections() error {
	for _, sec := range s.sections {
		str, err := sec.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/section", s.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *Body) tag() string {
	return "body"
}
