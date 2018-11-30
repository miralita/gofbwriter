package go_fbwriter

import "strings"

type author struct {
	firstName string
	middleName []string
	lastName string
	nickname []string
	homePage []string
	email []string
	id []string
	book *book
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

func (s *author) ToXml() (string, error) {
	if s.firstName != "" && s.lastName == "" {
		return "", makeError(ERR_EMPTY_FIRST_NAME, "Empty required field: author/first-name")
	} else if s.firstName == "" && s.lastName == "" {
		return "", makeError(ERR_EMPTY_FIELD, "Empty required field: author/nickname")
	}
	var b strings.Builder
	b.WriteString("<author>\n")
	b.WriteString(makeTag("first-name", s.firstName))
	b.WriteString(makeTagMulti("middle-name", s.middleName))
	b.WriteString(makeTag("last-name", s.lastName))
	b.WriteString(makeTagMulti("nickname", s.nickname))
	b.WriteString(makeTagMulti("home-page", s.homePage))
	b.WriteString(makeTagMulti("email", s.email))
	b.WriteString("</author>\n")
	return b.String(), nil
}