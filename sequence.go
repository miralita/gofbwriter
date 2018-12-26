package gofbwriter

type sequence struct{}

func (*sequence) ToXML() (string, error) {
	panic("implement me")
}

func (s *sequence) tag() string {
	return "sequence"
}
