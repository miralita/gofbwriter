package gofbwriter

import (
	"fmt"
	"strings"
)

//Information about a single author
type author struct {
	firstName  string
	middleName []string
	lastName   string
	nickname   []string
	homePage   []string
	email      []string
	id         []string
	tagName    string
	book       *book
}

func (s *author) setTagName(name string) {
	s.tagName = name
}

func (s *author) Email() []string {
	return s.email
}

func (s *author) AddEmail(email string) {
	if s.email == nil {
		s.email = []string{email}
	} else {
		s.email = append(s.email, email)
	}
}

func (s *author) HomePage() []string {
	return s.homePage
}

func (s *author) Nickname() []string {
	return s.nickname
}

func (s *author) MiddleName() []string {
	return s.middleName
}

func (s *author) LastName() string {
	return s.lastName
}

func (s *author) SetLastName(lastName string) {
	s.lastName = lastName
}

func (s *author) FirstName() string {
	return s.firstName
}

func (s *author) SetFirstName(firstName string) {
	s.firstName = firstName
}

func (s *author) AddMiddleName(name string) {
	if s.middleName == nil {
		s.middleName = []string{name}
	} else {
		s.middleName = append(s.middleName, name)
	}
}

func (s *author) AddNickname(name string) {
	if s.nickname == nil {
		s.nickname = []string{name}
	} else {
		s.nickname = append(s.nickname, name)
	}
}

func (s *author) AddHomepage(name string) {
	if s.nickname == nil {
		s.nickname = []string{name}
	} else {
		s.nickname = append(s.nickname, name)
	}
}

func (s *author) ToXML() (string, error) {
	if s.firstName != "" && s.lastName == "" {
		return "", makeError(ErrEmptyFirstName, "Empty required field: author/first-name")
	} else if s.firstName == "" && s.lastName == "" {
		return "", makeError(ErrEmptyField, "Empty required field: %s/nickname", s.tag())
	}
	var b strings.Builder

	fmt.Fprintf(&b, "<%s>\n", s.tag())
	b.WriteString(makeTag("first-name", s.firstName))
	b.WriteString(makeTagMulti("middle-name", s.middleName, true))
	b.WriteString(makeTag("last-name", s.lastName))
	b.WriteString(makeTagMulti("nickname", s.nickname, true))
	b.WriteString(makeTagMulti("home-page", s.homePage, true))
	b.WriteString(makeTagMulti("email", s.email, true))
	fmt.Fprintf(&b, "</%s>\n", s.tag())
	return b.String(), nil
}

func (s *author) tag() string {
	if s.tagName == "" {
		s.tagName = "author"
	}
	return s.tagName
}
