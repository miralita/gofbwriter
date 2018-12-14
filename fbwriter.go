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

func makeTag(tagName, tagValue string) string {
	if tagValue == "" {
		return ""
	}
	return fmt.Sprintf("<%s>%s</%s>\n", tagName, tagValue, tagName)
}

func makeTagMulti(tagName string, tagValue []string, sanitize bool) string {
	if tagValue == nil || len(tagValue) == 0 {
		return ""
	}
	ret := ""
	for _, s := range tagValue {
		if sanitize {
			s = sanitizeString(s)
		}
		ret += makeTag(tagName, s)
	}
	return ret
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

func makeAttribute(attrName, attrValue string) string {
	return fmt.Sprintf(`%s="%s"`, attrName, sanitizeString(attrValue))
}
