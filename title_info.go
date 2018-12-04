package gofbwriter

import "strings"

type titleInfo struct {
	genres    []Genre
	authors   []*author
	bookTitle string
	book      *book
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

func (s *titleInfo) AddAuthor(firstName, middleName, lastName string) *author {
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

func (s *titleInfo) ToXML() (string, error) {
	var b strings.Builder
	b.WriteString("<title-info>\n")
	if err := s.serializeGenres(&b); err != nil {
		return "", err
	}
	if err := s.serializeAuthors(&b); err != nil {
		return "", err
	}
	if err := s.serializeBookTitle(&b); err != nil {
		return "", err
	}

	b.WriteString("</title-info>")
	return b.String(), nil
}

func (s *titleInfo) serializeGenres(b *strings.Builder) error {
	if s.genres == nil || len(s.genres) == 0 {
		return makeError(ErrEmptyField, "Empty required field title-info/genre")
	}
	for _, g := range s.genres {
		b.WriteString(makeTag("genre", g.toString()))
	}
	return nil
}

func (s *titleInfo) serializeAuthors(b *strings.Builder) error {
	if s.authors == nil || len(s.authors) == 0 {
		return makeError(ErrEmptyField, "Empty required field title-info/author")
	}

	for _, a := range s.authors {
		xml, err := a.ToXML()
		if err != nil {
			return wrapError(err, ErrNestedEntity, "Can't make title-info/author")
		}
		b.WriteString(xml)
	}
	return nil
}

func (s *titleInfo) serializeBookTitle(b *strings.Builder) error {
	if s.bookTitle == "" {
		return makeError(ErrEmptyField, "Empty required field title-info/book-title")
	}
	b.WriteString(makeTag("book-title", s.bookTitle))
	return nil
}
