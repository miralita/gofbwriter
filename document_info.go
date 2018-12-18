package gofbwriter

type documentInfo struct{}

func (*documentInfo) ToXML() (string, error) {
	panic("implement me")
}
