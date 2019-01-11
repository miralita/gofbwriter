package gofbwriter

//Fb - common interface for book elements
type Fb interface {
	ToXML() (string, error)
	tag() string
	builder() *builder
}
