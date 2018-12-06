package gofbwriter

type poem struct {
	title *title
}

func (*poem) ToXML() (string, error) {
	panic("implement me")
}

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
