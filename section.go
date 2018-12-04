package gofbwriter

type poem struct {
	title *title
}

type annotation struct {
}

type epigraph struct {
	textAuthor string
	items      []*item
}

type section struct {
	title    *title
	epigraph *epigraph
	image    *image
}
