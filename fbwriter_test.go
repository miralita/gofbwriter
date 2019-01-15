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

func TestBookParsing(t *testing.T) {
	fb2 := gofbwriter.NewBook()
	fb2.SetTitle("Test book")
	fb2.SetLang("en")
	fb2.SetDocAuthor("fbwriter")
	fb2.SetGenre(gofbwriter.GenreAdventure)
	fb2.SetBookAuthor("Test", "Test", "Test")
	title := "Test section"
	body := `<p><b></p>
<p><em>Book One — Winds of Change Blow Over Pine Mist City</em></p>
<p><b></p>
<p>Chapter 1 - Chen Xi</p>
<p>Dusk fell upon Pine Mist City in the southern territory as the fiery sun set in the west.</p>
<p>For the thousandth time, Chen Xi pushed upon the door and entered the Zhang General Store.</p>
<p>The Zhang General Store was just an ordinary, medium-sized retail store within Pine Mist City that sold self-manufactured talisman-related products to keep itself afloat.</p>
<p>The merchandise that sold the most were first-grade and second-grade talismans. These were the foundation of the Zhang General Store’s survival. While its business wasn’t great, it benefitted from the small but steady stream of income and was barely able to establish itself within Pine Mist City.</p>
<p>“Talisman paper, talisman brush, and ink; it’s impossible to craft talismans without these three materials. It seems simple, but in reality it’s extremely complex. From today onwards, you will all learn how to differentiate between talisman papers, the utilization of the talisman brush, and the composition of the ink. Once you have a solid foundation, I’ll then instruct you in talisman crafting.”</p>
<p>Only now did Chen Xi realize that the store had once again recruited seven or eight talisman crafting apprentices with immature faces. Boss Zhang Dayong’s shriveled voice echoed within the general store.</p>
<p>“I’ll give the lot of you a month. If your skill doesn’t satisfy me after the month is over, then go home and play in the mud. You lot must remember that if you want to become a qualified talisman master, studying diligently and training hard is the only way to get there, as no one is able to easily succeed!”</p>
<p>The newly recruited talisman crafting apprentices had gazes that were filled with excitement and eagerness; they were itching to have a go at talisman crafting.</p>
<p>“Mmm, Chen Xi, you’re here.” Zhang Dayong looked over his shoulder to see Chen Xi and greeted him with a smile on his face.</p>
<p>“Uncle Zhang, these are the 30 Flamecloud Talismans for today.” Chen Xi pulled out a stack of azure talismans and passed them over.</p>
<p>Zhang Dayong waved his hand in dismissal. “There’s no rush. Since you’re here, then help me teach these little kids. These wages will be calculated separately. Hmmm, how about I pay you 3 spirit stones per hour?”</p>
<p>Chen Xi nodded after pondering for brief moment. “Alright!”</p>
`
	fb2.AddSection(title, body)
	t.Log(fb2.ToXML())
}
