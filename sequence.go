package gofbwriter

import "strconv"

//Book sequences
type sequence struct {
	b         *builder
	name      string
	number    int
	sequences []*sequence
}

// Add related sequence
func (s *sequence) AddSequence(item *sequence) {
	if s.sequences == nil {
		s.sequences = []*sequence{item}
	} else {
		s.sequences = append(s.sequences, item)
	}
}

// Create related sequence
func (s *sequence) CreateSequence() *sequence {
	seq := &sequence{}
	s.AddSequence(seq)
	return seq
}

func (s *sequence) Number() int {
	return s.number
}

func (s *sequence) SetNumber(number int) {
	s.number = number
}

func (s *sequence) Name() string {
	return s.name
}

func (s *sequence) SetName(name string) {
	s.name = name
}

func (s *sequence) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *sequence) ToXML() (string, error) {
	if s.name == "" {
		return "", makeError(ErrEmptyAttribute, "Empty required attribute %s/name", s.tag())
	}
	b := s.builder()
	b.openTagAttr(s.tag(), map[string]string{"name": s.name, "number": strconv.Itoa(s.number)}, false)
	b.closeTag(s.tag())
	if s.sequences != nil {
		for _, item := range s.sequences {
			str, err := item.ToXML()
			if err != nil {
				return "", wrapError(err, ErrNestedEntity, "Can't make %s/%s", s.tag(), item.tag())
			}
			b.WriteString(str)
		}
	}
	return b.String(), nil
}

func (s *sequence) tag() string {
	return "sequence"
}
