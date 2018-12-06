package gofbwriter

type fb interface {
	ToXML() (string, error)
}

type binary struct {
	ID          string
	ContentType string
	Data        []byte
}

type itype int

const (
	itypeP itype = iota
	itypeEmpty
)

type item struct {
	itype     itype
	itemValue fb
}
