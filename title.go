package gofbwriter

type title struct {
	items []fb
}

func (s *title) AddParagraph(str string) {
	s.addItem(&p{text: str})
}

func (s *title) addItem(i fb) {
	if s.items == nil {
		s.items = []fb{i}
	} else {
		s.items = append(s.items, i)
	}
}

func (s *title) AddEmptyline() {
	s.addItem(&empty{})
}
