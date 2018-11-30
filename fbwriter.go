package go_fbwriter

import "fmt"

func NewBook() *book {
	return &book{}
}

func makeTag(tag_name, tag_value string) string {
	if tag_value == "" {
		return ""
	}
	return fmt.Sprintf("<%s>%s</%s>\n", tag_name, tag_value, tag_name)
}

func makeTagMulti(tag_name string, tag_value []string) string {
	if tag_value == nil || len(tag_value) == 0 {
		return ""
	}
	ret := ""
	for _, s := range tag_value {
		ret += makeTag(tag_name, s)
	}
	return ret
}
