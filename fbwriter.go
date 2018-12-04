package gofbwriter

import (
	"fmt"
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

func makeTagMulti(tagName string, tagValue []string) string {
	if tagValue == nil || len(tagValue) == 0 {
		return ""
	}
	ret := ""
	for _, s := range tagValue {
		ret += makeTag(tagName, s)
	}
	return ret
}

func sanitizeString(str string) string {
	m := map[string]string{
		"<": "&lt;",
		">": "&gt;",
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
