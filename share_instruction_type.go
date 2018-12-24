package gofbwriter

//In-document instruction for generating output free and payed documents
type shareInstructionType struct {
	tagName string
}

func (*shareInstructionType) ToXML() (string, error) {
	panic("implement me")
}
