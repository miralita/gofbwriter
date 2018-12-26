package gofbwriter

import "strings"

type inlineImage struct {
	alt   string
	ctype string
	href  string
}

//An empty element with an image name as an attribute
type image struct {
	inlineImage
	title string
}

func (i *inlineImage) Href() string {
	return i.href
}

func (i *inlineImage) SetHref(href string) {
	i.href = href
}

func (i *inlineImage) Ctype() string {
	return i.ctype
}

func (i *inlineImage) SetCtype(ctype string) {
	i.ctype = ctype
}

func (i *inlineImage) Alt() string {
	return i.alt
}

func (i *inlineImage) SetAlt(alt string) {
	i.alt = alt
}

func (i *image) Title() string {
	return i.title
}

func (i *image) SetTitle(title string) {
	i.title = title
}

func (i *inlineImage) ToXML() (string, error) {
	var b strings.Builder
	b.WriteString("<image")
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
	b.WriteString(" />")
	return b.String(), nil
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
	b.WriteString(" />")
	return b.String(), nil
}

func (i *inlineImage) tag() string {
	return "image"
}

func (i *image) tag() string {
	return "image"
}
