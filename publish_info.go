package gofbwriter

import (
	"fmt"
)

//Information about some paper/outher published document, that was used as a source of this xml document
type publishInfo struct {
	b *builder
	//Original (paper) book name
	bookName string
	//Original (paper) book publisher
	publisher string
	//City where the original (paper) book was published
	city string
	//Year of the original (paper) publication
	year     int
	isbn     string
	sequence []*sequence
}

func (s *publishInfo) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *publishInfo) Sequence() []*sequence {
	return s.sequence
}

func (s *publishInfo) AddSequence(seq *sequence) {
	if s.sequence == nil {
		s.sequence = []*sequence{seq}
	} else {
		s.sequence = append(s.sequence, seq)
	}
}

func (s *publishInfo) CreateSequence() *sequence {
	seq := &sequence{}
	s.AddSequence(seq)
	return seq
}

func (s *publishInfo) Isbn() string {
	return s.isbn
}

func (s *publishInfo) SetIsbn(isbn string) {
	s.isbn = isbn
}

func (s *publishInfo) Year() int {
	return s.year
}

func (s *publishInfo) SetYear(year int) {
	s.year = year
}

func (s *publishInfo) City() string {
	return s.city
}

func (s *publishInfo) SetCity(city string) {
	s.city = city
}

func (s *publishInfo) Publisher() string {
	return s.publisher
}

func (s *publishInfo) SetPublisher(publisher string) {
	s.publisher = publisher
}

func (s *publishInfo) BookName() string {
	return s.bookName
}

func (s *publishInfo) SetBookName(bookName string) {
	s.bookName = bookName
}

func (s *publishInfo) ToXML() (string, error) {
	b := s.builder()
	b.openTag(s.tag())
	b.makeTag("book-name", sanitizeString(s.bookName))
	b.makeTag("publisher", sanitizeString(s.publisher))
	b.makeTag("city", sanitizeString(s.city))
	if s.year != 0 {
		b.makeTag("year", fmt.Sprintf("%d", s.year))
	}
	b.makeTag("isbn", sanitizeString(s.isbn))
	if s.sequence != nil {
		for _, seq := range s.sequence {
			str, err := seq.ToXML()
			if err != nil {
				return "", wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), seq.tag())
			}
			b.WriteString(str)
		}
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *publishInfo) tag() string {
	return "publish-info"
}
