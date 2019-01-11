package gofbwriter

//Author - Information about a single author
type Author struct {
	b          *builder
	firstName  string
	middleName []string
	lastName   string
	nickname   []string
	homePage   []string
	email      []string
	id         []string
	tagName    string
}

func (s *Author) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *Author) setTagName(name string) {
	s.tagName = name
}

//Emails - get list of author's emails
func (s *Author) Emails() []string {
	return s.email
}

//AddEmail - add new email to list of author's emails
func (s *Author) AddEmail(email string) {
	if s.email == nil {
		s.email = []string{email}
	} else {
		s.email = append(s.email, email)
	}
}

//HomePages - get list of author's home pages
func (s *Author) HomePages() []string {
	return s.homePage
}

//Nicknames - get list of author's nicknames
func (s *Author) Nicknames() []string {
	return s.nickname
}

//MiddleNames -  get list of author's middle names
func (s *Author) MiddleNames() []string {
	return s.middleName
}

//LastName - get author's last name
func (s *Author) LastName() string {
	return s.lastName
}

//SetLastName - set author's last name
func (s *Author) SetLastName(lastName string) {
	s.lastName = lastName
}

//FirstName - get author's first name
func (s *Author) FirstName() string {
	return s.firstName
}

//SetFirstName - set author's first name
func (s *Author) SetFirstName(firstName string) {
	s.firstName = firstName
}

//AddMiddleName - add new middle name to list of author's middle names
func (s *Author) AddMiddleName(name string) {
	if s.middleName == nil {
		s.middleName = []string{name}
	} else {
		s.middleName = append(s.middleName, name)
	}
}

//AddNickname - add new nickname to list of author's nicknames
func (s *Author) AddNickname(name string) {
	if s.nickname == nil {
		s.nickname = []string{name}
	} else {
		s.nickname = append(s.nickname, name)
	}
}

//AddHomepage - add new homepage to list of author's homepages
func (s *Author) AddHomepage(name string) {
	if s.nickname == nil {
		s.nickname = []string{name}
	} else {
		s.nickname = append(s.nickname, name)
	}
}

//ToXML - export to XML string
func (s *Author) ToXML() (string, error) {
	if s.firstName != "" && s.lastName == "" {
		return "", makeError(ErrEmptyFirstName, "Empty required field: author/first-name")
	} else if s.firstName == "" && s.lastName == "" && s.nickname == nil {
		return "", makeError(ErrEmptyField, "Empty required field: %s/nickname or %s/first-name + %s/last-name", s.tag(), s.tag(), s.tag())
	}
	b := s.builder()
	b.Reset()
	b.openTag(s.tag())
	b.makeTag("first-name", s.firstName)
	b.makeTags("middle-name", s.middleName, true)
	b.makeTag("last-name", s.lastName)
	b.makeTags("nickname", s.nickname, true)
	b.makeTags("home-page", s.homePage, true)
	b.makeTags("email", s.email, true)
	b.closeTag(s.tag())
	return b.String(), nil
}

func (s *Author) tag() string {
	if s.tagName == "" {
		s.tagName = "author"
	}
	return s.tagName
}
