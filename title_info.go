package gofbwriter

import (
	"strings"
	"time"
)

//Book (as a book opposite a document) description
type titleInfo struct {
	//Genre of this book, with the optional match percentage
	genres []Genre
	//Author(s) of this book
	authors []*author
	//Book title
	bookTitle string
	//Annotation for this book
	annotation *annotation
	//Any keywords for this book, intended for use in search engines
	keywords []string
	//Date this book was written, can be not exact, e.g. 1863-1867. If an optional attribute is present, then it should contain some computer-readable date from the interval for use by search and indexingengines
	date *date
	//Any coverpage items, currently only images
	coverpage *inlineImage
	//Book's language
	lang string
	//Book's source language if this is a translation
	srcLang string
	//Translators if this is a translation
	translators []*author
	//Any sequences this book might be part of
	sequences []*sequence
	book      *book
	tagName   string
}

func (s *titleInfo) CreateSequence() *sequence {
	seq := &sequence{}
	if s.sequences == nil {
		s.sequences = []*sequence{seq}
	} else {
		s.sequences = append(s.sequences, seq)
	}
	return seq
}

//Getter for sequences this book is part of
func (s *titleInfo) Sequences() []*sequence {
	return s.sequences
}

//Get translators list
func (s *titleInfo) Translators() []*author {
	return s.translators
}

//Creates a new translator, adds it to the list and returns reference
func (s *titleInfo) CreateTranslator() *author {
	tr := &author{tagName: "translator"}
	if s.translators == nil {
		s.translators = []*author{tr}
	} else {
		s.translators = append(s.translators, tr)
	}
	return tr
}

//Getter for source language
func (s *titleInfo) SrcLang() string {
	return s.srcLang
}

//Setter for source language
func (s *titleInfo) SetSrcLang(srcLang string) {
	s.srcLang = srcLang
}

//Getter for language
func (s *titleInfo) Lang() string {
	return s.lang
}

//Setter for language
func (s *titleInfo) SetLang(lang string) {
	s.lang = lang
}

func (s *titleInfo) Coverpage() *inlineImage {
	return s.coverpage
}

func (s *titleInfo) CreateCoverpage(href, ctype, alt string) *inlineImage {
	img := &inlineImage{}
	img.href = href
	img.ctype = ctype
	img.alt = alt
	s.coverpage = img
	return img
}

func (s *titleInfo) Date() *date {
	return s.date
}

func (s *titleInfo) SetDate(dt time.Time) {
	d := date(dt)
	s.date = &d
}

func (s *titleInfo) Annotation() *annotation {
	return s.annotation
}

func (s *titleInfo) CreateAnnotation() *annotation {
	s.annotation = &annotation{}
	return s.annotation
}

func (s *titleInfo) AddKeywords(keywords ...string) {
	if s.keywords == nil {
		s.keywords = keywords
	} else {
		s.keywords = append(s.keywords, keywords...)
	}
}

func (s *titleInfo) BookTitle() string {
	return s.bookTitle
}

func (s *titleInfo) SetBookTitle(bookTitle string) {
	s.bookTitle = bookTitle
}

func (s *titleInfo) Genres() []Genre {
	if s.genres == nil {
		s.genres = []Genre{}
	}
	return s.genres
}

func (s *titleInfo) AddGenre(genre Genre) []Genre {
	s.genres = append(s.Genres(), genre)
	return s.genres
}

func (s *titleInfo) Authors() []*author {
	if s.authors == nil {
		s.authors = []*author{}
	}
	return s.authors
}

func (s *titleInfo) CreateAuthor(firstName, middleName, lastName string) *author {
	author := &author{
		firstName:  firstName,
		middleName: []string{},
		lastName:   lastName,
		book:       s.book,
	}
	if middleName != "" {
		author.middleName = []string{middleName}
	}
	s.authors = append(s.Authors(), author)
	return author
}

func (s *titleInfo) openTag(b *strings.Builder) {
	if s.tagName == "" {
		s.tagName = "title-info"
	}
	b.WriteString("<")
	b.WriteString(s.tagName)
	b.WriteString(">\n")
}

func (s *titleInfo) closeTag(b *strings.Builder) {
	b.WriteString("</")
	b.WriteString(s.tagName)
	b.WriteString(">\n")
}

func (s *titleInfo) ToXML() (string, error) { // nolint: gocyclo
	var b strings.Builder

	if err := s.serializeGenres(&b); err != nil {
		return "", err
	}
	if err := s.serializeAuthors(&b); err != nil {
		return "", err
	}
	if err := s.serializeBookTitle(&b); err != nil {
		return "", err
	}
	if err := s.serializeAnnotation(&b); err != nil {
		return "", err
	}
	s.serializeKeywords(&b)
	if err := s.serializeDate(&b); err != nil {
		return "", err
	}
	if err := s.serializeCoverpage(&b); err != nil {
		return "", err
	}
	if err := s.serializeLang(&b); err != nil {
		return "", err
	}
	s.serializeSrcLang(&b)
	if err := s.serializeTranslators(&b); err != nil {
		return "", err
	}
	if err := s.serializeSequences(&b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (s *titleInfo) serializeSequences(b *strings.Builder) error {
	if s.sequences != nil {
		for _, tr := range s.sequences {
			str, err := tr.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/sequence", s.tagName)
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *titleInfo) serializeTranslators(b *strings.Builder) error {
	if s.translators != nil {
		for _, tr := range s.translators {
			str, err := tr.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make title-info/translator")
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *titleInfo) serializeSrcLang(b *strings.Builder) error {
	if s.lang != "" {
		b.WriteString(makeTag("src-lang", s.lang))
	}
	return nil
}

func (s *titleInfo) serializeLang(b *strings.Builder) error {
	if s.lang == "" {
		return makeError(ErrEmptyField, "Empty required %s/lang", s.tagName)
	}
	b.WriteString(makeTag("lang", s.lang))
	return nil
}

func (s *titleInfo) serializeCoverpage(b *strings.Builder) error {
	if s.coverpage != nil {
		str, err := s.coverpage.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/coverpage", s.tagName)
		}
		b.WriteString("<coverpage>")
		b.WriteString(str)
		b.WriteString("</coverpage>")
	}
	return nil
}

func (s *titleInfo) serializeDate(b *strings.Builder) error {
	if s.date != nil {
		str, err := s.date.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/date", s.tagName)
		}
		b.WriteString(str)
	}
	return nil
}

func (s *titleInfo) serializeKeywords(b *strings.Builder) error {
	if s.keywords != nil && len(s.keywords) > 0 {
		b.WriteString(makeTag("keywords", strings.Join(s.keywords, ",")))
	}
	return nil
}

func (s *titleInfo) serializeAnnotation(b *strings.Builder) error {
	if s.annotation != nil {
		str, err := s.annotation.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/annotation", s.tagName)
		}
		b.WriteString(str)
	}
	return nil
}

func (s *titleInfo) serializeGenres(b *strings.Builder) error {
	if s.genres == nil || len(s.genres) == 0 {
		return makeError(ErrEmptyField, "Empty required field %s/genre", s.tagName)
	}
	for _, g := range s.genres {
		b.WriteString(makeTag("genre", g.toString()))
	}
	return nil
}

func (s *titleInfo) serializeAuthors(b *strings.Builder) error {
	if s.authors == nil || len(s.authors) == 0 {
		return makeError(ErrEmptyField, "Empty required field %s/author", s.tagName)
	}

	for _, a := range s.authors {
		xml, err := a.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/author", s.tagName)
		}
		b.WriteString(xml)
	}
	return nil
}

func (s *titleInfo) serializeBookTitle(b *strings.Builder) error {
	if s.bookTitle == "" {
		return makeError(ErrEmptyField, "Empty required field %s/book-title", s.tagName)
	}
	b.WriteString(makeTag("book-title", s.bookTitle))
	return nil
}
