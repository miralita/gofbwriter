package gofbwriter

import (
	"fmt"
	"strings"
)

type description struct {
	b *builder
	//Generic information about the book
	titleInfo *titleInfo
	//Generic information about the original book (for translations)
	srcTitleInfo *titleInfo
	//Information about this particular (xml) document
	documentInfo *documentInfo
	//Information about some paper/outher published document, that was used as a source of this xml document
	publishInfo *publishInfo
	//Any other information about the book/document that didnt fit in the above groups
	customInfo []*customInfo
	//Describes, how the document should be presented to end-user, what parts are free, what parts should be sold and what price should be used
	output []*shareInstructionType
	book   *book
}

func (s *description) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *description) Output() []*shareInstructionType {
	return s.output
}

func (s *description) CreateOutput() (*shareInstructionType, error) {
	if s.output != nil && len(s.output) >= 2 {
		return nil, makeError(ErrToManyItems, "Can't create new description/output: max occurs = 2")
	}
	sit := &shareInstructionType{tagName: "output"}
	if s.output == nil {
		s.output = []*shareInstructionType{sit}
	} else {
		s.output = append(s.output, sit)
	}
	return sit, nil
}

func (s *description) CustomInfo() []*customInfo {
	return s.customInfo
}

func (s *description) AddCustomInfo(infoType, value string) {
	info := &customInfo{info: value, infoType: infoType}
	if s.customInfo == nil {
		s.customInfo = []*customInfo{info}
	} else {
		s.customInfo = append(s.customInfo, info)
	}
}

func (s *description) PublishInfo() *publishInfo {
	return s.publishInfo
}

func (s *description) CreatePublishInfo() *publishInfo {
	s.publishInfo = &publishInfo{}
	return s.publishInfo
}

func (s *description) DocumentInfo() *documentInfo {
	return s.documentInfo
}

func (s *description) CreateDocumentInfo() *documentInfo {
	s.documentInfo = &documentInfo{}
	return s.documentInfo
}

func (s *description) SrcTitleInfo() *titleInfo {
	return s.srcTitleInfo
}

func (s *description) CreateSrcTitleInfo() *titleInfo {
	s.srcTitleInfo = &titleInfo{tagName: "src-title-info"}
	return s.srcTitleInfo
}

func (s *description) TitleInfo() *titleInfo {
	return s.titleInfo
}

func (s *description) CreateTitleInfo() *titleInfo {
	s.titleInfo = &titleInfo{}
	return s.titleInfo
}

func (s *description) ToXML() (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "<%s>\n", s.tag())
	if err := s.serializeTitleInfo(&b); err != nil {
		return "", err
	}
	if err := s.serializeSrcTitleInfo(&b); err != nil {
		return "", err
	}
	if err := s.serializeDocumentInfo(&b); err != nil {
		return "", err
	}
	if err := s.serializePublishInfo(&b); err != nil {
		return "", err
	}
	if err := s.serializeCustomInfo(&b); err != nil {
		return "", err
	}
	if err := s.serializeOutput(&b); err != nil {
		return "", err
	}
	fmt.Fprintf(&b, "</%s>\n", s.tag())
	return b.String(), nil
}

func (s *description) serializeOutput(b *strings.Builder) error {
	if s.output != nil {
		for _, o := range s.output {
			str, err := o.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/output", s.tag())
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *description) serializeCustomInfo(b *strings.Builder) error {
	if s.customInfo != nil {
		for _, i := range s.customInfo {
			str, err := i.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/custom-info", s.tag())
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *description) serializePublishInfo(b *strings.Builder) error {
	if s.publishInfo == nil {
		return makeError(ErrEmptyField, "Empty required %s/publish-info", s.tag())
	}
	str, err := s.documentInfo.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/publish-info", s.tag())
	}
	b.WriteString(str)
	return nil
}

func (s *description) serializeDocumentInfo(b *strings.Builder) error {
	if s.documentInfo == nil {
		return makeError(ErrEmptyField, "Empty required %s/document-info", s.tag())
	}
	str, err := s.documentInfo.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/document-info", s.tag())
	}
	b.WriteString(str)
	return nil
}

func (s *description) serializeSrcTitleInfo(b *strings.Builder) error {
	if s.srcTitleInfo != nil {
		str, err := s.srcTitleInfo.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/src-title-info", s.tag())
		}
		b.WriteString(str)
	}
	return nil
}

func (s *description) serializeTitleInfo(b *strings.Builder) error {
	if s.titleInfo == nil {
		return makeError(ErrEmptyField, "Empty required %s/title-info", s.tag())
	}
	str, err := s.titleInfo.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/title-info", s.tag())
	}
	b.WriteString(str)
	return nil
}

func (s *description) tag() string {
	return "description"
}
