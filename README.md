# gofbwriter

Simple library for generating FB2 documents

# Example

## Use wrappers

```go
package main
import (
	"log"
	"fmt"
	"github.com/miralita/gofbwriter"
)
func main() {
	fb2 := gofbwriter.NewBook()
    fb2.SetTitle("Test book")
    fb2.SetLang("en")
    fb2.SetDocAuthor("fbwriter")
    fb2.SetGenre(gofbwriter.GenreAdventure)
    fb2.SetBookAuthor("Test", "Test", "Test")
    title := "Test section"
    body := `<p>First paragraph</p><p>Second paragraph</p>`
    _, err := fb2.AddSection(title, body)
    if err != nil {
    	log.Fatalf("Can't add section to book: %s", err)
    }
    xml, err := fb2.ToXML()
    if err != nil {
    	log.Fatalf("Can't export book to FB2 XML: %s", err)
    }
    fmt.Println(xml)
    err = fb2.Save("path/to/book.fb2")
    if err != nil {
    	log.Fatalf("Can't save book to file: %s", err)
    }
}
```

## Use FB2 structure directly

```go
package main
import (
	"log"
	"fmt"
	"github.com/miralita/gofbwriter"
)
func main() {
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
    err := sec.AddParagraph("Test")
    if err != nil {
        log.Fatalf("Can't add paragraph to section: %s", err)
    }


    err = book.Save("tmp/test.fb2")
    if err != nil {
    	log.Fatalf("Can't save book to file: %s", err)
    }
}
```
