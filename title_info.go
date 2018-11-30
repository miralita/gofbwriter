package go_fbwriter

import "strings"

type titleInfo struct {
	genres []Genre
	authors []*author
	bookTitle string
	book *book
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
		firstName: firstName,
		middleName: []string{},
		lastName: lastName,
		book: s.book,
	}
	if middleName != "" {
		author.middleName = []string{middleName}
	}
	s.authors = append(s.Authors(), author)
	return author
}

func (s *titleInfo) ToXml() (string, error) {
	if s.genres == nil || len(s.genres) == 0 {
		return "", makeError(ERR_EMPTY_FIELD, "Empty required field title-info/genre")
	}
	var b strings.Builder
	b.WriteString("<title-info>\n")
	for _, g := range s.genres {
		b.WriteString(makeTag("genre", g.ToString()))
	}
	b.WriteString("</title-info>")
	return b.String(), nil
}


