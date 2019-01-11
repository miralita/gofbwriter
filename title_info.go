package gofbwriter

import (
	"strings"
	"time"
)

//TitleInfo - Book (as a book opposite a document) description
type TitleInfo struct {
	b *builder
	//Genre of this book, with the optional match percentage
	genres []Genre
	//Author(s) of this book
	authors []*Author
	//Book title
	bookTitle string
	//Annotation for this book
	annotation *Annotation
	//Any keywords for this book, intended for use in search engines
	keywords []string
	//Date this book was written, can be not exact, e.g. 1863-1867. If an optional attribute is present, then it should contain some computer-readable date from the interval for use by search and indexingengines
	date *Date
	//Any coverpage items, currently only images
	coverpage *InlineImage
	//Book's language
	lang string
	//Book's source language if this is a translation
	srcLang string
	//Translators if this is a translation
	translators []*Author
	//Any sequences this book might be part of
	sequences []*Sequence
	tagName   string
}

func (s *TitleInfo) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//CreateSequence - create new Sequence, add to list of sequences and return
func (s *TitleInfo) CreateSequence() *Sequence {
	seq := &Sequence{}
	if s.sequences == nil {
		s.sequences = []*Sequence{seq}
	} else {
		s.sequences = append(s.sequences, seq)
	}
	return seq
}

//Sequences - Getter for sequences this book is part of
func (s *TitleInfo) Sequences() []*Sequence {
	return s.sequences
}

//Translators - Get translators list
func (s *TitleInfo) Translators() []*Author {
	return s.translators
}

//CreateTranslator - create a new translator, add it to the list and return reference
func (s *TitleInfo) CreateTranslator() *Author {
	tr := &Author{tagName: "translator"}
	if s.translators == nil {
		s.translators = []*Author{tr}
	} else {
		s.translators = append(s.translators, tr)
	}
	return tr
}

//SrcLang - get source language (for translation)
func (s *TitleInfo) SrcLang() string {
	return s.srcLang
}

//SetSrcLang - set source language
func (s *TitleInfo) SetSrcLang(srcLang string) {
	s.srcLang = srcLang
}

//Lang - get language
func (s *TitleInfo) Lang() string {
	return s.lang
}

//SetLang - set language
func (s *TitleInfo) SetLang(lang string) {
	s.lang = lang
}

//Coverpage - get cover image
func (s *TitleInfo) Coverpage() *InlineImage {
	return s.coverpage
}

//CreateCoverpage - create and return new image
func (s *TitleInfo) CreateCoverpage(href, ctype, alt string) *InlineImage {
	img := &InlineImage{}
	img.href = href
	img.ctype = ctype
	img.alt = alt
	s.coverpage = img
	return img
}

//Date - Date this book was written, can be not exact, e.g. 1863-1867. If an optional attribute is present, then it should contain some computer-readable date from the interval for use by search and indexingengines
func (s *TitleInfo) Date() *Date {
	return s.date
}

//SetDate - set date this book was written
func (s *TitleInfo) SetDate(dt time.Time) {
	d := Date(dt)
	s.date = &d
}

//Annotation - get annotation
func (s *TitleInfo) Annotation() *Annotation {
	return s.annotation
}

//CreateAnnotation - create and return new Annotation; old value will be dropped
func (s *TitleInfo) CreateAnnotation() *Annotation {
	s.annotation = &Annotation{}
	return s.annotation
}

//AddKeywords - add keywords to list
func (s *TitleInfo) AddKeywords(keywords ...string) {
	if s.keywords == nil {
		s.keywords = keywords
	} else {
		s.keywords = append(s.keywords, keywords...)
	}
}

//BookTitle - get book title
func (s *TitleInfo) BookTitle() string {
	return s.bookTitle
}

//SetBookTitle - set book title
func (s *TitleInfo) SetBookTitle(bookTitle string) {
	s.bookTitle = bookTitle
}

//Genres - get list of genres
func (s *TitleInfo) Genres() []Genre {
	if s.genres == nil {
		s.genres = []Genre{}
	}
	return s.genres
}

//AddGenre - add new genre to list and return modified list
func (s *TitleInfo) AddGenre(genre Genre) []Genre {
	s.genres = append(s.Genres(), genre)
	return s.genres
}

//Authors - get list of book authors
func (s *TitleInfo) Authors() []*Author {
	if s.authors == nil {
		s.authors = []*Author{}
	}
	return s.authors
}

//CreateAuthor - create new Author from given field, add it to list of author and return
func (s *TitleInfo) CreateAuthor(firstName, middleName, lastName string) *Author {
	author := &Author{
		firstName:  firstName,
		middleName: []string{},
		lastName:   lastName,
	}
	if middleName != "" {
		author.middleName = []string{middleName}
	}
	s.authors = append(s.Authors(), author)
	return author
}

//ToXML - export to XML string
func (s *TitleInfo) ToXML() (string, error) { // nolint: gocyclo
	b := s.builder()
	b.Reset()
	b.openTag(s.tag())

	if err := s.serializeGenres(); err != nil {
		return "", err
	}
	if err := s.serializeAuthors(); err != nil {
		return "", err
	}
	if err := s.serializeBookTitle(); err != nil {
		return "", err
	}
	if err := s.serializeAnnotation(); err != nil {
		return "", err
	}
	_ = s.serializeKeywords()
	if err := s.serializeDate(); err != nil {
		return "", err
	}
	if err := s.serializeCoverpage(); err != nil {
		return "", err
	}
	if err := s.serializeLang(); err != nil {
		return "", err
	}
	_ = s.serializeSrcLang()
	if err := s.serializeTranslators(); err != nil {
		return "", err
	}
	if err := s.serializeSequences(); err != nil {
		return "", err
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *TitleInfo) serializeSequences() error {
	if s.sequences != nil {
		for _, tr := range s.sequences {
			str, err := tr.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/sequence", s.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *TitleInfo) serializeTranslators() error {
	if s.translators != nil {
		for _, tr := range s.translators {
			str, err := tr.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make title-info/translator")
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *TitleInfo) serializeSrcLang() error {
	s.builder().makeTag("src-lang", s.lang)
	return nil
}

func (s *TitleInfo) serializeLang() error {
	if s.lang == "" {
		return makeError(ErrEmptyField, "Empty required %s/lang", s.tag())
	}
	s.builder().makeTag("lang", s.lang)
	return nil
}

func (s *TitleInfo) serializeCoverpage() error {
	if s.coverpage != nil {
		str, err := s.coverpage.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/coverpage", s.tag())
		}
		s.builder().makeTag("coverpage", str)
	}
	return nil
}

func (s *TitleInfo) serializeDate() error {
	if s.date != nil {
		str, err := s.date.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/date", s.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *TitleInfo) serializeKeywords() error {
	if s.keywords != nil && len(s.keywords) > 0 {
		s.builder().makeTag("keywords", strings.Join(s.keywords, ","))
	}
	return nil
}

func (s *TitleInfo) serializeAnnotation() error {
	if s.annotation != nil {
		str, err := s.annotation.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/annotation", s.tag())
		}
		s.builder().WriteString(str)
	}
	return nil
}

func (s *TitleInfo) serializeGenres() error {
	if s.genres == nil || len(s.genres) == 0 {
		return makeError(ErrEmptyField, "Empty required field %s/genre", s.tag())
	}
	for _, g := range s.genres {
		s.builder().makeTag("genre", g.toString())
	}
	return nil
}

func (s *TitleInfo) serializeAuthors() error {
	if s.authors == nil || len(s.authors) == 0 {
		return makeError(ErrEmptyField, "Empty required field %s/author", s.tag())
	}

	for _, a := range s.authors {
		xml, err := a.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make %s/author", s.tag())
		}
		s.builder().WriteString(xml)
	}
	return nil
}

func (s *TitleInfo) serializeBookTitle() error {
	if s.bookTitle == "" {
		return makeError(ErrEmptyField, "Empty required field %s/book-title", s.tag())
	}
	s.builder().makeTag("book-title", s.bookTitle)
	return nil
}

func (s *TitleInfo) tag() string {
	return "title-info"
}
