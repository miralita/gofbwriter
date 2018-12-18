package gofbwriter

import "strings"

type description struct {
	//Generic information about the book
	titleInfo *titleInfo
	//Generic information about the original book (for translations)
	srcTitleInfo *titleInfo
	//Information about this particular (xml) document
	documentInfo *documentInfo
	book         *book
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
	b.WriteString("<description>\n")
	if err := s.serializeTitleInfo(&b); err != nil {
		return "", err
	}
	if err := s.serializeSrcTitleInfo(&b); err != nil {
		return "", err
	}
	if err := s.serializeDocumentInfo(&b); err != nil {
		return "", err
	}
	b.WriteString("</description>\n")
	return b.String(), nil
}

func (s *description) serializeDocumentInfo(b *strings.Builder) error {
	if s.documentInfo == nil {
		return makeError(ErrEmptyField, "Empty required description/document-info")
	}
	str, err := s.documentInfo.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make description/document-info")
	}
	b.WriteString(str)
	return nil
}

func (s *description) serializeSrcTitleInfo(b *strings.Builder) error {
	if s.srcTitleInfo != nil {
		str, err := s.srcTitleInfo.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make description/src-title-info")
		}
		b.WriteString(str)
	}
	return nil
}

func (s *description) serializeTitleInfo(b *strings.Builder) error {
	if s.titleInfo == nil {
		return makeError(ErrEmptyField, "Empty required description/title-info")
	}
	str, err := s.titleInfo.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make description/title-info")
	}
	b.WriteString(str)
	return nil
}
