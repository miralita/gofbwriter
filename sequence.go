package gofbwriter

import "strconv"

//Sequence - Book sequences
type Sequence struct {
	b         *builder
	name      string
	number    int
	sequences []*Sequence
}

//AddSequence - Add related sequence
func (s *Sequence) AddSequence(item *Sequence) {
	if s.sequences == nil {
		s.sequences = []*Sequence{item}
	} else {
		s.sequences = append(s.sequences, item)
	}
}

//CreateSequence - Create new related sequence, add it to list and return
func (s *Sequence) CreateSequence() *Sequence {
	seq := &Sequence{}
	s.AddSequence(seq)
	return seq
}

//Number - get number of the book in sequence
func (s *Sequence) Number() int {
	return s.number
}

//SetNumber - set number of the book in sequence
func (s *Sequence) SetNumber(number int) {
	s.number = number
}

//Name - get name of sequence
func (s *Sequence) Name() string {
	return s.name
}

//SetName - set name of sequence
func (s *Sequence) SetName(name string) {
	s.name = name
}

func (s *Sequence) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

//ToXML - export to XML string
func (s *Sequence) ToXML() (string, error) {
	if s.name == "" {
		return "", makeError(ErrEmptyAttribute, "Empty required attribute %s/name", s.tag())
	}
	b := s.builder()
	b.Reset()
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

func (s *Sequence) tag() string {
	return "sequence"
}
