package gofbwriter

type inlineImage struct {
	b     *builder
	alt   string
	ctype string
	href  string
}

func (i *inlineImage) builder() *builder {
	if i.b == nil {
		i.b = &builder{}
	}
	return i.b
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
	i.builder().Reset()
	i.builder().makeTagAttr(i.tag(), "", map[string]string{"ctype": i.ctype, "alt": i.alt, "href": i.href}, false)
	return i.builder().String(), nil
}

func (i *image) ToXML() (string, error) {
	i.builder().Reset()
	i.builder().makeTagAttr(i.tag(), "", map[string]string{"ctype": i.ctype, "alt": i.alt, "href": i.href, "title": i.title}, false)
	return i.builder().String(), nil
}

func (i *inlineImage) tag() string {
	return "image"
}

func (i *image) tag() string {
	return "image"
}
