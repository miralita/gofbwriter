package gofbwriter

//InlineImage - simple image type
type InlineImage struct {
	b     *builder
	alt   string
	ctype string
	href  string
}

func (i *InlineImage) builder() *builder {
	if i.b == nil {
		i.b = &builder{}
	}
	return i.b
}

//Image - an empty element with an image name as an attribute
type Image struct {
	InlineImage
	title string
}

//Href - get href attribute
func (i *InlineImage) Href() string {
	return i.href
}

//SetHref - set href attribute
func (i *InlineImage) SetHref(href string) {
	i.href = href
}

//Ctype - get type attribute
func (i *InlineImage) Ctype() string {
	return i.ctype
}

//SetCtype - set type attribute
func (i *InlineImage) SetCtype(ctype string) {
	i.ctype = ctype
}

//Alt - get alt attribute
func (i *InlineImage) Alt() string {
	return i.alt
}

//SetAlt - set alt attribute
func (i *InlineImage) SetAlt(alt string) {
	i.alt = alt
}

//Title - get title attribute
func (i *Image) Title() string {
	return i.title
}

//SetTitle - set title attribute
func (i *Image) SetTitle(title string) {
	i.title = title
}

//ToXML - export to XML string
func (i *InlineImage) ToXML() (string, error) {
	i.builder().Reset()
	i.builder().makeTagAttr(i.tag(), "", map[string]string{"ctype": i.ctype, "alt": i.alt, "href": i.href}, false)
	return i.builder().String(), nil
}

//ToXML - export to XML string
func (i *Image) ToXML() (string, error) {
	i.builder().Reset()
	i.builder().makeTagAttr(i.tag(), "", map[string]string{"ctype": i.ctype, "alt": i.alt, "href": i.href, "title": i.title}, false)
	return i.builder().String(), nil
}

func (i *InlineImage) tag() string {
	return "image" //nolint:goconst
}

func (i *Image) tag() string {
	return "image" //nolint:goconst
}
