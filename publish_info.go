package gofbwriter

import (
	"fmt"
)

//PublishInfo - Information about some paper/outher published document, that was used as a source of this xml document
type PublishInfo struct {
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
	sequence []*Sequence
}

func (s *PublishInfo) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Sequences - get list of book sequences
func (s *PublishInfo) Sequences() []*Sequence {
	return s.sequence
}

//AddSequence - add existing sequence to list
func (s *PublishInfo) AddSequence(seq *Sequence) {
	if s.sequence == nil {
		s.sequence = []*Sequence{seq}
	} else {
		s.sequence = append(s.sequence, seq)
	}
}

//CreateSequence - create new Sequence, add it to list and return
func (s *PublishInfo) CreateSequence() *Sequence {
	seq := &Sequence{}
	s.AddSequence(seq)
	return seq
}

//ISBN - get book ISBN
func (s *PublishInfo) ISBN() string {
	return s.isbn
}

//SetISBN - set book ISBN
func (s *PublishInfo) SetISBN(isbn string) {
	s.isbn = isbn
}

//Year - get year of the original (paper) publication
func (s *PublishInfo) Year() int {
	return s.year
}

//SetYear - set year of the original (paper) publication
func (s *PublishInfo) SetYear(year int) {
	s.year = year
}

//City - get city where the original (paper) book was published
func (s *PublishInfo) City() string {
	return s.city
}

//SetCity - set city where the original (paper) book was published
func (s *PublishInfo) SetCity(city string) {
	s.city = city
}

//Publisher - get original (paper) book publisher
func (s *PublishInfo) Publisher() string {
	return s.publisher
}

//SetPublisher - set original (paper) book publisher
func (s *PublishInfo) SetPublisher(publisher string) {
	s.publisher = publisher
}

//BookName - get original (paper) book name
func (s *PublishInfo) BookName() string {
	return s.bookName
}

//SetBookName - original (paper) book name
func (s *PublishInfo) SetBookName(bookName string) {
	s.bookName = bookName
}

//ToXML - export to XML string
func (s *PublishInfo) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
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

func (s *PublishInfo) tag() string {
	return "publish-info"
}
