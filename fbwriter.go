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

func (b *builder) makeTag(tagName, tagValue string) {
	if tagValue == "" {
		return
	}
	fmt.Fprintf(b, "<%s>%s</%s>\n", tagName, tagValue, tagName)
}

func (b *builder) makeTags(tagName string, tagValue []string, sanitize bool) {
	if tagValue == nil || len(tagValue) == 0 {
		return
	}
	for _, s := range tagValue {
		if sanitize {
			s = sanitizeString(s)
		}
		b.makeTag(tagName, s)
	}
}

func (b *builder) openTag(tagName string, attrs map[string]string, sanitize bool) {
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
}

func (b *builder) closeTag(tagName string) {
	fmt.Fprintf(b, "</%s>\n", tagName)
}

func sanitizeString(str string) string {
	re := regexp.MustCompile(`<(/?[A-z0-9-]+).*?>`)
	str = re.ReplaceAllString(str, "<$1>")
	str = re.ReplaceAllStringFunc(str, func(s string) string {
		goodTags := []string{"b", "i", "strong", "del", "em", "pre", "small", "sub", "sup", "u"}
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

func (b *builder) makeAttribute(attrName, attrValue string) {
	fmt.Fprintf(b, `%s="%s"`, attrName, sanitizeString(attrValue))
}
