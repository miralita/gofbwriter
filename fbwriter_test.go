package gofbwriter_test

import (
	"github.com/miralita/go-fbwriter"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	book := gofbwriter.NewBook()
	t.Log(book)

	var b strings.Builder
	b.WriteString("test\n")
	addToBuilder(&b, "value1\n")
	addToBuilder(&b, "value1\n")
	b.WriteString("test")
	t.Log(b.String())
}

func addToBuilder(b *strings.Builder, value string) {
	b.WriteString(value)
}
