package gofbwriter

import (
	"fmt"
	"regexp"
	"strings"
)

//NewBook - creates new empty book
func NewBook() *book { // nolint: golint
	return &book{}
}

type builder struct {
	strings.Builder
}

func (b *builder) makeTag(tagName, tagValue string) *builder {
	if tagValue == "" {
		return b
	}
	fmt.Fprintf(b, "<%s>%s</%s>\n", tagName, tagValue, tagName)
	return b
}

func (b *builder) makeTags(tagName string, tagValue []string, sanitize bool) *builder {
	if tagValue == nil || len(tagValue) == 0 {
		return b
	}
	for _, s := range tagValue {
		if sanitize {
			s = sanitizeString(s)
		}
		b.makeTag(tagName, s)
	}
	return b
}

func (b *builder) openTagAttr(tagName string, attrs map[string]string, sanitize bool) *builder {
	fmt.Fprintf(b, "<%s", tagName)
	if attrs != nil {
		for k, v := range attrs {
			if v != "" {
				if sanitize {
					v = sanitizeString(v)
				}
				fmt.Fprintf(b, " %s=\"%s\"", k, v)
			}
		}
	}
	fmt.Fprint(b, ">\n")
	return b
}

func (b *builder) makeTagAttr(tagName string, value string, attrs map[string]string, sanitize bool) *builder {
	fmt.Fprintf(b, "<%s", tagName)
	if attrs != nil {
		for k, v := range attrs {
			if v != "" {
				if sanitize {
					v = sanitizeString(v)
				}
				fmt.Fprintf(b, " %s=\"%s\"", k, v)
			}
		}
	}
	if value != "" {
		fmt.Fprint(b, value)
		fmt.Fprint(b, ">\n")
	} else {
		fmt.Fprint(b, " />\n")
	}
	return b
}

func (b *builder) openTag(tagName string) *builder {
	fmt.Fprintf(b, "<%s>\n", tagName)
	return b
}

func (b *builder) closeTag(tagName string) *builder {
	fmt.Fprintf(b, "</%s>\n", tagName)
	return b
}

func sanitizeString(str string) string {
	re := regexp.MustCompile(`<(/?[A-z0-9-]+).*?>`)
	str = re.ReplaceAllString(str, "<$1>")
	str = re.ReplaceAllStringFunc(str, func(s string) string {
		goodTags := []string{"b", "i", "strong", "del", "em", "pre", "small", "sub", "sup", "u", "strikethrough", "emphasis"}
		for _, t := range goodTags {
			if strings.HasPrefix(s, "<"+t) || strings.HasPrefix(s, "</"+t) {
				return s
			}
		}
		return ""
	})
	m := map[string]string{
		`"`: "&quot;",
		"'": "&apos;",
	}
	for k, v := range m {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}

func (b *builder) makeAttribute(attrName, attrValue string) *builder {
	fmt.Fprintf(b, `%s="%s"`, attrName, sanitizeString(attrValue))
	return b
}
