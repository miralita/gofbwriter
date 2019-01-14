package gofbwriter

import (
	"fmt"
	"io/ioutil"
)

//Fb2 - FictionBook v2 struct
type Fb2 struct {
	b *builder
	//This element contains an arbitrary stylesheet that is intepreted by a some processing programs, e.g. text/css stylesheets can be used by XSLT stylesheets to generate better looking html
	stylesheets []*StyleSheet
	description *Description
	//Main content of the book, multiple bodies are used for additional information, like footnotes, that do not appear in the main book flow. The first body is presented to the reader by default, and content in the other bodies should be accessible by hyperlinks. Name attribute should describe the meaning of this body, this is optional for the main body.
	body *Body
	//Body for footnotes, content is mostly similar to base type and may (!) be rendered in the pure environment "as is". Advanced reader should treat section[2]/section as endnotes, all other stuff as footnotes
	notes []*Body
	//Any binary data that is required for the presentation of this book in base64 format. Currently only images are used.
	binary []*Binary
}

func (s *Fb2) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//Binary - get list of binary
func (s *Fb2) Binary() []*Binary {
	return s.binary
}

//CreateBinary - creates new binary struct and adds it to list of binary, returns created struct
func (s *Fb2) CreateBinary(id, contentType string, data []byte) *Binary {
	b := &Binary{id: id, contentType: contentType, data: data}
	if s.binary == nil {
		s.binary = []*Binary{b}
	} else {
		s.binary = append(s.binary, b)
	}
	return b
}

//Notes - get list of notes
func (s *Fb2) Notes() []*Body {
	return s.notes
}

//CreateNote - create new note and add it to list of notes, return created struct
func (s *Fb2) CreateNote(name string) *Body {
	b := &Body{name: name}
	if s.notes == nil {
		s.notes = []*Body{b}
	} else {
		s.notes = append(s.notes, b)
	}
	return b
}

//Body - get Body struct
func (s *Fb2) Body() *Body {
	return s.body
}

//CreateBody - create new Body, old Body will be dropped
func (s *Fb2) CreateBody() *Body {
	s.body = &Body{}
	return s.body
}

//CreateStylesheet - create new stylesheet and add it to list of stylesheets, return created struct
func (s *Fb2) CreateStylesheet(ctype, data string) *StyleSheet {
	if ctype == "" {
		ctype = "text/css"
	}
	st := &StyleSheet{&builder{}, ctype, data}
	if s.stylesheets == nil {
		s.stylesheets = []*StyleSheet{st}
	} else {
		s.stylesheets = append(s.stylesheets, st)
	}
	return st
}

//Description - get description
func (s *Fb2) Description() *Description {
	if s.description == nil {
		s.description = &Description{}
	}
	return s.description
}

//CreateDescription - creates new description, old description will be dropped
func (s *Fb2) CreateDescription() *Description {
	s.description = &Description{}
	return s.description
}

//ToXML - convert book to xml string
func (s *Fb2) ToXML() (string, error) {
	b := s.builder()
	b.Reset()
	fmt.Fprintf(b, `<?xml version="1.0" encoding="UTF-8"?>
<%s xmlns="http://www.gribuser.ru/xml/fictionbook/2.0" xmlns:xlink="http://www.w3.org/1999/xlink">
`, s.tag())
	if err := s.makeStylesheets(); err != nil {
		return "", err
	}
	if err := s.makeDescription(); err != nil {
		return "", err
	}
	if err := s.makeBody(); err != nil {
		return "", err
	}
	if err := s.makeNotes(); err != nil {
		return "", err
	}
	if err := s.makeBinary(); err != nil {
		return "", err
	}
	b.closeTag(s.tag())
	return b.String(), nil
}

//Save - save book to file
func (s *Fb2) Save(filename string) error {
	str, err := s.ToXML()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, []byte(str), 0644)
}

func (s *Fb2) makeBinary() error {
	if s.binary != nil && len(s.binary) > 0 {
		for _, bin := range s.binary {
			str, err := bin.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/binary", s.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

//SetTitle - save book's title
func (s *Fb2) SetTitle(title string) *Fb2 {
	s.titleInfo().SetBookTitle(title)
	return s
}

func (s *Fb2) titleInfo() *TitleInfo {
	descr := s.Description()
	if descr == nil {
		descr = s.CreateDescription()
		descr.CreatePublishInfo()
	}
	titleInfo := descr.TitleInfo()
	if titleInfo == nil {
		titleInfo = descr.CreateTitleInfo()
	}
	return titleInfo
}

//SetLang - set book language
func (s *Fb2) SetLang(lang string) *Fb2 {
	s.titleInfo().SetLang(lang)
	return s
}

//SetBookAuthor - save author
func (s *Fb2) SetBookAuthor(firstName, middleName, lastName string) *Fb2 {
	s.titleInfo().CreateAuthor(firstName, middleName, lastName)
	return s
}

//SetDocAuthor - save author of current document
func (s *Fb2) SetDocAuthor(nickname string) *Fb2 {
	descr := s.Description()
	if descr == nil {
		descr = s.CreateDescription()
		descr.CreatePublishInfo()
	}
	docInfo := descr.DocumentInfo()
	if docInfo == nil {
		docInfo = descr.CreateDocumentInfo()
	}
	docInfo.CreateAuthor().AddNickname(nickname)
	return s
}

//AddSection - add entire text block with title to book
func (s *Fb2) AddSection(title, body string) (*Section, error) {
	bd := s.Body()
	if bd == nil {
		bd = s.CreateBody()
	}
	sec := bd.CreateSection()
	if title != "" {
		t := sec.CreateTitle()
		t.AddParagraph(title)
	}
	data := prepareSection(body)
	for _, str := range data {
		sec.AddParagraph(str)
	}
	return sec, nil
}

//SetGenre - save book's genre
func (s *Fb2) SetGenre(genre Genre) {
	s.titleInfo().AddGenre(genre)
}

func (s *Fb2) makeNotes() error {
	if s.notes != nil && len(s.notes) > 0 {
		for _, n := range s.notes {
			str, err := n.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make %s/body (notes)", s.tag())
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *Fb2) makeBody() error {
	if s.body == nil {
		return makeError(ErrEmptyField, "Empty required %s/body", s.tag())
	}
	str, err := s.body.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/body", s.tag())
	}
	s.builder().WriteString(str)
	return nil
}

func (s *Fb2) makeDescription() error {
	if s.description == nil {
		return makeError(ErrEmptyField, "Empty required %s/description", s.tag())
	}
	str, err := s.description.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make %s/description", s.tag())
	}
	s.builder().WriteString(str)
	return nil
}

func (s *Fb2) makeStylesheets() error {
	if s.stylesheets != nil && len(s.stylesheets) > 0 {
		for _, st := range s.stylesheets {
			str, err := st.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make FictionBook/stylesheet")
			}
			s.builder().WriteString(str)
		}
	}
	return nil
}

func (s *Fb2) tag() string {
	return "FictionBook"
}
