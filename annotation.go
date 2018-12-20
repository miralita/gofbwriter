package gofbwriter

type annotation struct {
	tagName string
}

func (*annotation) ToXML() (string, error) {
	panic("implement me")
}

func (s *annotation) AddParagraph(p string) {
	panic("implement me")
}
