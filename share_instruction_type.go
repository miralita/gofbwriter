package gofbwriter

type shareInstructionType struct {
	tagName string
}

func (*shareInstructionType) ToXML() (string, error) {
	panic("implement me")
}
