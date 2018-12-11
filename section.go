package gofbwriter

type annotation struct {
}

type section struct {
	title    *title
	epigraph *epigraph
	image    *image
}

func (*section) ToXML() (string, error) {
	panic("implement me")
}
