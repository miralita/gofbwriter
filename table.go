package gofbwriter

type table struct{}

func (*table) ToXML() (string, error) {
	panic("implement me")
}

func (s *table) tag() string {
	return "table"
}
