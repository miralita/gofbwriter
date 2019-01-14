package gofbwriter_test

import (
	"github.com/miralita/gofbwriter"
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

func TestNewBook(t *testing.T) {
	book := gofbwriter.NewBook()
	descr := book.CreateDescription()
	titleInfo := descr.CreateTitleInfo()
	titleInfo.AddGenre(gofbwriter.GenreAdventure)
	titleInfo.CreateAuthor("Иван", "Иванович", "Иванов")
	titleInfo.SetBookTitle("Тестовая книга")
	titleInfo.SetLang("ru")

	docInfo := descr.CreateDocumentInfo()
	docInfo.CreateAuthor().AddNickname("miralita")

	descr.CreatePublishInfo()

	body := book.CreateBody()
	sec := body.CreateSection()
	sec.AddParagraph("Test")

	//t.Log(author)

	str, err := book.ToXML()
	t.Log(str, err)
	t.Log(book.Save("tmp/test.fb2"))
}
