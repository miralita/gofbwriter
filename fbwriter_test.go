package go_fbwriter_test

import (
	"go-fbwriter"
	"testing"
)

func TestInit(t *testing.T) {
	book := go_fbwriter.NewBook()
	t.Log(book)
}
