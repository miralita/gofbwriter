package gofbwriter

type description struct {
	titleInfo *titleInfo
	book      *book
}

func (s *description) TitleInfo() *titleInfo {
	if s.titleInfo == nil {
		s.titleInfo = &titleInfo{book: s.book}
	}
	return s.titleInfo
}
