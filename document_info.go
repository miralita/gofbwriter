package gofbwriter

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

//DocumentInfo - Information about this particular (xml) document
type DocumentInfo struct {
	b *builder
	//Author(s) of this particular document
	authors []*Author
	//Any software used in preparation of this document, in free format
	programUsed string
	//Date this document was created, same guidelines as in the &lt;title-info&gt; section apply
	date *Date
	//Source URL if this document is a conversion of some other (online) document
	srcUrls []string
	//URL of the original (online) document, if this is a conversion
	srcOcr string
	//this is a unique identifier for a document. this must not change
	id *uuid.UUID
	//Document version, in free format, should be incremented if the document is changed and re-released to the public
	version float32
	//Short description for all changes made to this document, like "Added missing chapter 6", in free form.
	history []*Annotation
	//Owner of the fb2 document copyrights
	publishers []*Author
}

func (s *DocumentInfo) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Publishers - Owner of the fb2 document copyrights
func (s *DocumentInfo) Publishers() []*Author {
	return s.publishers
}

//CreatePublisher - create new Author, add to list of publishers and return
func (s *DocumentInfo) CreatePublisher() *Author {
	a := &Author{tagName: "publisher"}
	s.AddPublisher(a)
	return a
}

//AddPublisher - add existing author to publishers
func (s *DocumentInfo) AddPublisher(publisher *Author) {
	if s.publishers == nil {
		s.publishers = []*Author{publisher}
	} else {
		s.publishers = append(s.publishers, publisher)
	}
}

//History - Short description for all changes made to this document, like "Added missing chapter 6", in free form.
func (s *DocumentInfo) History() []*Annotation {
	return s.history
}

//AddHistory - add single record to history; return created Annotation struct
func (s *DocumentInfo) AddHistory(descr string) *Annotation {
	d := &Annotation{tagName: "history"}
	d.AddParagraph(descr)
	if s.history == nil {
		s.history = []*Annotation{d}
	} else {
		s.history = append(s.history, d)
	}
	return d
}

//Version - Document version, in free format, should be incremented if the document is changed and re-released to the public
func (s *DocumentInfo) Version() float32 {
	return s.version
}

//SetVersion - set document's version
func (s *DocumentInfo) SetVersion(version float32) {
	s.version = version
}

//ID - this is a unique identifier for a document. this must not change
func (s *DocumentInfo) ID() uuid.UUID {
	return *s.id
}

//SetID - set ID
func (s *DocumentInfo) SetID(id uuid.UUID) {
	s.id = &id
}

//SrcOcr - URL of the original (online) document, if this is a conversion
func (s *DocumentInfo) SrcOcr() string {
	return s.srcOcr
}

//SetSrcOcr - set url of original document
func (s *DocumentInfo) SetSrcOcr(srcOcr string) {
	s.srcOcr = srcOcr
}

//SrcUrls - Source URL if this document is a conversion of some other (online) document
func (s *DocumentInfo) SrcUrls() []string {
	return s.srcUrls
}

//AddSrcURL - add new src url to list
func (s *DocumentInfo) AddSrcURL(url string) {
	if s.srcUrls == nil {
		s.srcUrls = []string{url}
	} else {
		s.srcUrls = append(s.srcUrls, url)
	}
}

//SetDate - set date this document was created
func (s *DocumentInfo) SetDate(dt time.Time) {
	d := Date(dt)
	s.date = &d
}

//Date - Date this document was created, same guidelines as in the title-info; section apply
func (s *DocumentInfo) Date() *Date {
	return s.date
}

//ProgramUsed - Any software used in preparation of this document, in free format
func (s *DocumentInfo) ProgramUsed() string {
	return s.programUsed
}

//SetProgramUsed - set name of software used for creating of this document
func (s *DocumentInfo) SetProgramUsed(programUsed string) {
	s.programUsed = programUsed
}

//Authors - Author(s) of this particular document
func (s *DocumentInfo) Authors() []*Author {
	return s.authors
}

//CreateAuthor - create new Author, add it to list of authors and return
func (s *DocumentInfo) CreateAuthor() *Author {
	a := &Author{}
	s.AddAuthor(a)
	return a
}

//AddAuthor - add existing Author to list of authors
func (s *DocumentInfo) AddAuthor(docAuthor *Author) {
	if s.authors == nil {
		s.authors = []*Author{docAuthor}
	} else {
		s.authors = append(s.authors, docAuthor)
	}
}

//ToXML - export to XML string
func (s *DocumentInfo) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
	b.openTag(s.tag())
	if err := s.serializeAuthors(); err != nil {
		return "", err
	}
	if s.programUsed != "" {
		b.makeTag("program-used", sanitizeString(s.programUsed))
	}
	if err := s.serializeDate(); err != nil {
		return "", err
	}
	if s.srcOcr != "" {
		b.makeTag("src-ocr", sanitizeString(s.srcOcr))
	}
	s.serializeID()
	if s.version == 0 {
		s.version = 1.0
	}
	b.makeTag("version", fmt.Sprintf("%0.4f", s.version))
	if err := s.serializeHistory(); err != nil {
		return "", err
	}
	if err := s.serializePublisher(); err != nil {
		return "", err
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *DocumentInfo) serializePublisher() error {
	if s.publishers != nil {
		for _, h := range s.publishers {
			str, err := h.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/publisher", s.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *DocumentInfo) serializeHistory() error {
	if s.history != nil {
		for _, h := range s.history {
			str, err := h.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/history", s.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *DocumentInfo) serializeID() {
	if s.id == nil {
		u := uuid.NewV4()
		s.id = &u
	}
	s.builder().makeTag("id", s.id.String())
}

func (s *DocumentInfo) serializeSrcUrls() error {
	if s.srcUrls != nil {
		for _, url := range s.srcUrls {
			s.builder().makeTag("src-url", sanitizeString(url))
		}
	}
	return nil
}

func (s *DocumentInfo) serializeDate() error {
	if s.date == nil {
		s.SetDate(time.Now())
	}
	str, err := s.date.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/date", s.tag())
	}
	s.builder().WriteString(str)
	return nil
}

func (s *DocumentInfo) serializeAuthors() error {
	if s.authors == nil || len(s.authors) == 0 {
		return makeError(ErrEmptyField, "Empty required %s/author", s.tag())
	}
	for _, a := range s.authors {
		str, err := a.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/author", s.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *DocumentInfo) tag() string {
	return "document-info"
}
