package gofbwriter

type fb interface {
	ToXML() (string, error)
	tag() string
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
