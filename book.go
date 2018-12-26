package gofbwriter

import "strings"

//Root element
type book struct {
	//This element contains an arbitrary stylesheet that is intepreted by a some processing programs, e.g. text/css stylesheets can be used by XSLT stylesheets to generate better looking html
	stylesheets []*stylesheet
	description *description
	//Main content of the book, multiple bodies are used for additional information, like footnotes, that do not appear in the main book flow. The first body is presented to the reader by default, and content in the other bodies should be accessible by hyperlinks. Name attribute should describe the meaning of this body, this is optional for the main body.
	body *body
	//Body for footnotes, content is mostly similar to base type and may (!) be rendered in the pure environment "as is". Advanced reader should treat section[2]/section as endnotes, all other stuff as footnotes
	notes []*body
	//Any binary data that is required for the presentation of this book in base64 format. Currently only images are used.
	binary []*binary
}

func (s *book) Binary() []*binary {
	return s.binary
}

func (s *book) CreateBinary(id, contentType string, data []byte) *binary {
	b := &binary{id: id, contentType: contentType, data: data}
	if s.binary == nil {
		s.binary = []*binary{b}
	} else {
		s.binary = append(s.binary, b)
	}
	return b
}

func (s *book) Notes() []*body {
	return s.notes
}

func (s *book) CreateNote(name string) *body {
	b := &body{name: name}
	if s.notes == nil {
		s.notes = []*body{b}
	} else {
		s.notes = append(s.notes, b)
	}
	return b
}

func (s *book) Body() *body {
	return s.body
}

func (s *book) CreateBody() *body {
	s.body = &body{}
	return s.body
}

func (s *book) CreateStylesheet(ctype, data string) (*stylesheet, error) {
	if ctype == "" {
		ctype = "text/css"
	}
	st := &stylesheet{ctype, data, s}
	if s.stylesheets == nil {
		s.stylesheets = []*stylesheet{st}
	} else {
		s.stylesheets = append(s.stylesheets, st)
	}
	return st, nil
}

func (s *book) Description() *description {
	if s.description == nil {
		s.description = &description{book: s}
	}
	return s.description
}

func (s *book) CreateDescription() *description {
	s.description = &description{}
	return s.description
}

func (s *book) ToXML() (string, error) {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>
<FictionBook xmlns="http://www.gribuser.ru/xml/fictionbook/2.0" xmlns:xlink="http://www.w3.org/1999/xlink">
`)
	if err := s.makeStylesheets(&b); err != nil {
		return "", err
	}
	if err := s.makeDescription(&b); err != nil {
		return "", err
	}
	if err := s.makeBody(&b); err != nil {
		return "", err
	}
	if err := s.makeNotes(&b); err != nil {
		return "", err
	}
	if err := s.makeBinary(&b); err != nil {
		return "", err
	}
	b.WriteString("</FictionBook>\n")
	return b.String(), nil
}

func (s *book) makeBinary(b *strings.Builder) error {
	if s.binary != nil && len(s.binary) > 0 {
		for _, bin := range s.binary {
			str, err := bin.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make FictionBook/binary")
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *book) makeNotes(b *strings.Builder) error {
	if s.notes != nil && len(s.notes) > 0 {
		for _, n := range s.notes {
			str, err := n.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make FictionBook/body (notes)")
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *book) makeBody(b *strings.Builder) error {
	if s.body == nil {
		return makeError(ErrEmptyField, "Empty required FictionBook/body")
	}
	str, err := s.body.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make FictionBook/body")
	}
	b.WriteString(str)
	return nil
}

func (s *book) makeDescription(b *strings.Builder) error {
	if s.description == nil {
		return makeError(ErrEmptyField, "Empty required FictionBook/description")
	}
	str, err := s.description.ToXML()
	if err != nil {
		return wrapError(err, ErrNestedEntity, "Can't make FictionBook/description")
	}
	b.WriteString(str)
	return nil
}

func (s *book) makeStylesheets(b *strings.Builder) error {
	if s.stylesheets != nil && len(s.stylesheets) > 0 {
		for _, st := range s.stylesheets {
			str, err := st.ToXML()
			if err != nil {
				return wrapError(err, ErrNestedEntity, "Can't make FictionBook/stylesheet")
			}
			b.WriteString(str)
		}
	}
	return nil
}

func (s *book) tag() string {
	return "FictionBook"
}
