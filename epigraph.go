package gofbwriter

import (
	"fmt"
	"strings"
)

//An epigraph
/*<xs:complexType name="epigraphType">
  <xs:annotation>
    <xs:documentation>An epigraph</xs:documentation>
  </xs:annotation>
  <xs:sequence>
    <xs:choice minOccurs="0" maxOccurs="unbounded">
      <xs:element name="p" type="pType"/>
      <xs:element name="poem" type="poemType"/>
      <xs:element name="cite" type="citeType"/>
      <xs:element name="empty-line"/>
    </xs:choice>
    <xs:element name="text-author" type="pType" minOccurs="0" maxOccurs="unbounded"/>
  </xs:sequence>
  <xs:attribute name="id" type="xs:ID" use="optional"/>
</xs:complexType>*/
type epigraph struct {
	b           *builder
	textAuthors []string
	items       []fb
}

func (s *epigraph) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *epigraph) TextAuthors() []string {
	return s.textAuthors
}

func (s *epigraph) AddAuthor(author string) {
	if s.textAuthors == nil {
		s.textAuthors = []string{author}
	} else {
		s.textAuthors = append(s.textAuthors, author)
	}
}

func (s *epigraph) CreateParagraph(str string) *p {
	p := &p{text: str}
	_ = s.AppendItem(p)
	return p
}

func (s *epigraph) CreateEmptyLine() {
	_ = s.AppendItem(&empty{})
}

func (s *epigraph) CreatePoem() *poem {
	p := &poem{}
	_ = s.AppendItem(p)
	return p
}

func (s *epigraph) CreateCite() *cite {
	p := &cite{}
	_ = s.AppendItem(p)
	return p
}

func (s *epigraph) AppendItem(item fb) error {
	pass := false
	if _, ok := item.(*p); ok {
		pass = true
	} else if _, ok := item.(*poem); ok {
		pass = true
	} else if _, ok := item.(*cite); ok {
		pass = true
	} else if _, ok := item.(*empty); ok {
		pass = true
	}
	if !pass {
		return makeError(ErrUnsupportedNestedItem, "Can't use type %T in epigraph", item)
	}
	if s.items == nil {
		s.items = []fb{item}
	} else {
		s.items = append(s.items, item)
	}
	return nil
}

func (s *epigraph) ToXML() (string, error) {
	if (s.items == nil || len(s.items) == 0) && (s.textAuthors == nil || len(s.textAuthors) == 0) {
		return fmt.Sprintf("<%s />\n", s.tag()), nil
	}
	var b strings.Builder
	fmt.Fprintf(&b, "<%s>\n", s.tag())
	if s.items != nil {
		for _, i := range s.items {
			str, err := i.ToXML()
			if err != nil {
				return "", wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), i.tag())
			}
			b.WriteString(str)
		}
	}
	if s.textAuthors != nil {
		for _, i := range s.textAuthors {
			b.WriteString(makeTag("text-author", sanitizeString(i)))
		}
	}
	fmt.Fprintf(&b, "</%s>\n", s.tag())
	return b.String(), nil
}

func (s *epigraph) tag() string {
	return "epigraph"
}
