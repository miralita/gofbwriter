package gofbwriter

import "strings"

type image struct {
	title string
	alt   string
	ctype string
	href  string
}

func (i *image) Href() string {
	return i.href
}

func (i *image) SetHref(href string) {
	i.href = href
}

func (i *image) Ctype() string {
	return i.ctype
}

func (i *image) SetCtype(ctype string) {
	i.ctype = ctype
}

func (i *image) Alt() string {
	return i.alt
}

func (i *image) SetAlt(alt string) {
	i.alt = alt
}

func (i *image) Title() string {
	return i.title
}

func (i *image) SetTitle(title string) {
	i.title = title
}

func (i *image) ToXML() (string, error) {
	var b strings.Builder
	b.WriteString("<image")
	if i.title != "" {
		b.WriteString(" ")
		b.WriteString(makeAttribute("title", i.title))
	}
	if i.ctype != "" {
		b.WriteString(" ")
		b.WriteString(makeAttribute("ctype", i.ctype))
	}
	if i.alt != "" {
		b.WriteString(" ")
		b.WriteString(makeAttribute("alt", i.alt))
	}
	if i.href != "" {
		b.WriteString(" ")
		b.WriteString(makeAttribute("href", i.href))
	}
	return b.String(), nil
}
