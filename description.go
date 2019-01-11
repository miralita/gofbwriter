package gofbwriter

//Description - description of the book
type Description struct {
	b *builder
	//Generic information about the book
	titleInfo *TitleInfo
	//Generic information about the original book (for translations)
	srcTitleInfo *TitleInfo
	//Information about this particular (xml) document
	documentInfo *DocumentInfo
	//Information about some paper/outher published document, that was used as a source of this xml document
	publishInfo *PublishInfo
	//Any other information about the book/document that didnt fit in the above groups
	customInfo []*CustomInfo
	//Describes, how the document should be presented to end-user, what parts are free, what parts should be sold and what price should be used
	output []*ShareInstruction
}

func (s *Description) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Output - Describes, how the document should be presented to end-user, what parts are free, what parts should be sold and what price should be used
func (s *Description) Output() []*ShareInstruction {
	return s.output
}

//CreateOutput - create new ShareInstruction, add it to list of outputs and return
func (s *Description) CreateOutput() (*ShareInstruction, error) {
	if s.output != nil && len(s.output) >= 2 {
		return nil, makeError(ErrToManyItems, "Can't create new description/output: max occurs = 2")
	}
	sit := &ShareInstruction{tagName: "output"}
	if s.output == nil {
		s.output = []*ShareInstruction{sit}
	} else {
		s.output = append(s.output, sit)
	}
	return sit, nil
}

//CustomInfo - Any other information about the book/document that didnt fit in the above groups
func (s *Description) CustomInfo() []*CustomInfo {
	return s.customInfo
}

//AddCustomInfo - add new CustomInfo to list
func (s *Description) AddCustomInfo(infoType, value string) {
	info := &CustomInfo{info: value, infoType: infoType}
	if s.customInfo == nil {
		s.customInfo = []*CustomInfo{info}
	} else {
		s.customInfo = append(s.customInfo, info)
	}
}

//PublishInfo - Information about some paper/outher published document, that was used as a source of this xml document
func (s *Description) PublishInfo() *PublishInfo {
	return s.publishInfo
}

//CreatePublishInfo - create and return new PublishInfo. Old PublishInfo will be dropped
func (s *Description) CreatePublishInfo() *PublishInfo {
	s.publishInfo = &PublishInfo{}
	return s.publishInfo
}

//DocumentInfo - Information about this particular (xml) document
func (s *Description) DocumentInfo() *DocumentInfo {
	return s.documentInfo
}

//CreateDocumentInfo - create and return new DocumentInfo. Old DocumentInfo will be dropped
func (s *Description) CreateDocumentInfo() *DocumentInfo {
	s.documentInfo = &DocumentInfo{}
	return s.documentInfo
}

//SrcTitleInfo - Generic information about the original book (for translations)
func (s *Description) SrcTitleInfo() *TitleInfo {
	return s.srcTitleInfo
}

//CreateSrcTitleInfo - create and return new src-title-info. Old value will be dropped
func (s *Description) CreateSrcTitleInfo() *TitleInfo {
	s.srcTitleInfo = &TitleInfo{tagName: "src-title-info"}
	return s.srcTitleInfo
}

//TitleInfo - Generic information about the book
func (s *Description) TitleInfo() *TitleInfo {
	return s.titleInfo
}

//CreateTitleInfo - create and return new title-info. Old value will be dropped
func (s *Description) CreateTitleInfo() *TitleInfo {
	s.titleInfo = &TitleInfo{}
	return s.titleInfo
}

//ToXML - export to XML string
func (s *Description) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
	b.openTag(s.tag())
	if err := s.serializeTitleInfo(); err != nil {
		return "", err
	}
	if err := s.serializeSrcTitleInfo(); err != nil {
		return "", err
	}
	if err := s.serializeDocumentInfo(); err != nil {
		return "", err
	}
	if err := s.serializePublishInfo(); err != nil {
		return "", err
	}
	if err := s.serializeCustomInfo(); err != nil {
		return "", err
	}
	if err := s.serializeOutput(); err != nil {
		return "", err
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *Description) serializeOutput() error {
	if s.output != nil {
		for _, o := range s.output {
			str, err := o.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/output", s.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *Description) serializeCustomInfo() error {
	if s.customInfo != nil {
		for _, i := range s.customInfo {
			str, err := i.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/custom-info", s.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *Description) serializePublishInfo() error {
	if s.publishInfo == nil {
		return makeError(ErrEmptyField, "Empty required %s/publish-info", s.tag())
	}
	str, err := s.publishInfo.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/publish-info", s.tag())
	}
	s.builder().WriteString(str)
	return nil
}

func (s *Description) serializeDocumentInfo() error {
	if s.documentInfo == nil {
		return makeError(ErrEmptyField, "Empty required %s/document-info", s.tag())
	}
	str, err := s.documentInfo.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/document-info", s.tag())
	}
	s.builder().WriteString(str)
	return nil
}

func (s *Description) serializeSrcTitleInfo() error {
	if s.srcTitleInfo != nil {
		str, err := s.srcTitleInfo.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/src-title-info", s.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *Description) serializeTitleInfo() error {
	if s.titleInfo == nil {
		return makeError(ErrEmptyField, "Empty required %s/title-info", s.tag())
	}
	str, err := s.titleInfo.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/title-info", s.tag())
	}
	s.builder().WriteString(str)
	return nil
}

func (s *Description) tag() string {
	return "description"
}
