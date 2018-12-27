package gofbwriter

import (
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

//Information about this particular (xml) document
type documentInfo struct {
	//Author(s) of this particular document
	authors []*author
	//Any software used in preparation of this document, in free format
	programUsed string
	//Date this document was created, same guidelines as in the &lt;title-info&gt; section apply
	date *date
	//Source URL if this document is a conversion of some other (online) document
	srcUrls []string
	//URL of the original (online) document, if this is a conversion
	srcOcr string
	//this is a unique identifier for a document. this must not change
	id *uuid.UUID
	//Document version, in free format, should be incremented if the document is changed and re-released to the public
	version float32
	//Short description for all changes made to this document, like "Added missing chapter 6", in free form.
	history []*annotation
	//Owner of the fb2 document copyrights
	publishers []*author
}

func (s *documentInfo) Publishers() []*author {
	return s.publishers
}

func (s *documentInfo) CreatePublisher() *author {
	a := &author{tagName: "publisher"}
	s.AddPublisher(a)
	return a
}

func (s *documentInfo) AddPublisher(publisher *author) {
	if s.publishers == nil {
		s.publishers = []*author{publisher}
	} else {
		s.publishers = append(s.publishers, publisher)
	}
}

func (s *documentInfo) History() []*annotation {
	return s.history
}

func (s *documentInfo) AddHistory(descr string) *annotation {
	d := &annotation{tagName: "history"}
	d.AddParagraph(descr)
	if s.history == nil {
		s.history = []*annotation{d}
	} else {
		s.history = append(s.history, d)
	}
	return d
}

func (s *documentInfo) Version() float32 {
	return s.version
}

func (s *documentInfo) SetVersion(version float32) {
	s.version = version
}

func (s *documentInfo) ID() uuid.UUID {
	return *s.id
}

func (s *documentInfo) SetID(id uuid.UUID) {
	s.id = &id
}

func (s *documentInfo) SrcOcr() string {
	return s.srcOcr
}

func (s *documentInfo) SetSrcOcr(srcOcr string) {
	s.srcOcr = srcOcr
}

func (s *documentInfo) SrcUrls() []string {
	return s.srcUrls
}

func (s *documentInfo) AddSrcURL(url string) {
	if s.srcUrls == nil {
		s.srcUrls = []string{url}
	} else {
		s.srcUrls = append(s.srcUrls, url)
	}
}

func (s *documentInfo) SetDate(dt time.Time) {
	d := date(dt)
	s.date = &d
}

func (s *documentInfo) Date() *date {
	return s.date
}

func (s *documentInfo) ProgramUsed() string {
	return s.programUsed
}

func (s *documentInfo) SetProgramUsed(programUsed string) {
	s.programUsed = programUsed
}

func (s *documentInfo) Authors() []*author {
	return s.authors
}

func (s *documentInfo) CreateAuthor() *author {
	a := &author{}
	s.AddAuthor(a)
	return a
}

func (s *documentInfo) AddAuthor(docAuthor *author) {
	if s.authors == nil {
		s.authors = []*author{docAuthor}
	} else {
		s.authors = append(s.authors, docAuthor)
	}
}

func (s *documentInfo) ToXML() (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "<%s>\n", s.tag())
	if err := s.serializeAuthors(&b); err != nil {
		return "", err
	}
	if s.programUsed != "" {
		b.WriteString(makeTag("program-used", sanitizeString(s.programUsed)))
	}
	if err := s.serializeDate(&b); err != nil {
		return "", err
	}
	if s.srcOcr != "" {
		b.WriteString(makeTag("src-ocr", sanitizeString(s.srcOcr)))
	}
	s.serializeID(&b)
	if s.version == 0 {
		s.version = 1.0
	}
	b.WriteString(makeTag("version", fmt.Sprintf("%0.4f", s.version)))
	if err := s.serializeHistory(&b); err != nil {
		return "", err
	}
	if err := s.serializePublisher(&b); err != nil {
		return "", err
	}
	fmt.Fprintf(&b, "</%s>\n", s.tag())
	return b.String(), nil
}

func (s *documentInfo) serializePublisher(b *strings.Builder) error {
	if s.publishers != nil {
		for _, h := range s.publishers {
			str, err := h.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/publisher", s.tag())
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *documentInfo) serializeHistory(b *strings.Builder) error {
	if s.history != nil {
		for _, h := range s.history {
			str, err := h.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/history", s.tag())
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *documentInfo) serializeID(b *strings.Builder) {
	if s.id == nil {
		u := uuid.NewV4()
		s.id = &u
	}
	b.WriteString(makeTag("id", s.id.String()))
}

func (s *documentInfo) serializeSrcUrls(b *strings.Builder) error {
	if s.srcUrls != nil {
		for _, url := range s.srcUrls {
			b.WriteString(makeTag("src-url", sanitizeString(url)))
		}
	}
	return nil
}

func (s *documentInfo) serializeDate(b *strings.Builder) error {
	if s.date == nil {
		s.SetDate(time.Now())
	}
	str, err := s.date.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/date", s.tag())
	}
	b.WriteString(str)
	return nil
}

func (s *documentInfo) serializeAuthors(b *strings.Builder) error {
	if s.authors == nil || len(s.authors) == 0 {
		return makeError(ErrEmptyField, "Empty required %s/author", s.tag())
	}
	for _, a := range s.authors {
		str, err := a.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/author", s.tag())
		}
		b.WriteString(str)
	}
	return nil
}

func (s *documentInfo) tag() string {
	return "document-info"
}
