package gofbwriter

type sequence struct{}

func (*sequence) ToXML() (string, error) {
	panic("implement me")
}
