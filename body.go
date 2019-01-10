package gofbwriter

//Main content of the book, multiple bodies are used for additional information, like footnotes, that do not appear in the main book flow (extended from this class). The first body is presented to the reader by default, and content in the other bodies should be accessible by hyperlinks.
type body struct {
	b         *builder
	image     *image      //Image to be displayed at the top of this section
	title     *title      //A fancy title for the entire book, should be used if the simple text version in &lt;description&gt; is not adequate, e.g. the book title has multiple paragraphs and/or character styles
	epigraphs []*epigraph //Epigraph(s) for the entire book, if any
	sections  []*section
	name      string
	book      *book
}

func (s *body) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *body) Sections() []*section {
	return s.sections
}

func (s *body) AddSection(sec *section) {
	if s.sections == nil {
		s.sections = []*section{sec}
	} else {
		s.sections = append(s.sections, sec)
	}
}

func (s *body) CreateSection() *section {
	sec := &section{}
	s.AddSection(sec)
	return sec
}

func (s *body) Epigraphs() []*epigraph {
	return s.epigraphs
}

func (s *body) AddEpigraph(ep *epigraph) {
	if s.epigraphs == nil {
		s.epigraphs = []*epigraph{ep}
	} else {
		s.epigraphs = append(s.epigraphs, ep)
	}
}

func (s *body) CreateEpigraph() *epigraph {
	ep := &epigraph{}
	s.AddEpigraph(ep)
	return ep
}

func (s *body) Title() *title {
	return s.title
}

func (s *body) SetTitle(title *title) {
	s.title = title
}

func (s *body) CreateTitle() *title {
	s.title = &title{}
	return s.title
}

func (s *body) Image() *image {
	return s.image
}

func (s *body) CreateImage() *image {
	img := &image{}
	s.image = img
	return img
}

func (s *body) SetImage(image *image) {
	s.image = image
}

func (s *body) ToXML() (string, error) {
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

func (s *body) serializeImage() error {
	if s.image != nil {
		i, err := s.image.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/image", s.tag())
		}
		s.builder().WriteString(i)
	}
	return nil
}

func (s *body) serializeTitle() error {
	if s.title != nil {
		t, err := s.title.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/title", s.tag())
		}
		s.builder().WriteString(t)
	}
	return nil
}

func (s *body) serializeEpigraphs() error {
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

func (s *body) serializeSections() error {
	for _, sec := range s.sections {
		str, err := sec.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/section", s.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *body) tag() string {
	return "body"
}
