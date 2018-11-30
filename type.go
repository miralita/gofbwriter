package go_fbwriter

type fbwriter interface {
	ToXml() (string, error)
}

type body struct {
	book *book
}

type binary struct {
	Id string
	ContentType string
	Data []byte
}



