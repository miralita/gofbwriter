package gofbwriter

//DocGenerationInstructionType - List of instructions to process sections (allow|deny|require)
type DocGenerationInstructionType int

//Value - text value for doc_generation_instruction_type
func (s DocGenerationInstructionType) Value() (string, error) {
	val := ""
	switch s {
	case InstructionTypeRequire:
		val = "require"
	case InstructionTypeAllow:
		val = "allow"
	case InstructionTypeDeny:
		val = "deny"
	default:
		return "", makeError(ErrUnknownEnumValue, "Can't use value %s for DocGenerationInstructionType enumeration", s.String())
	}
	return val, nil
}

//go:generate stringer -type=DocGenerationInstructionType

const (
	//InstructionTypeRequire - Require
	InstructionTypeRequire DocGenerationInstructionType = iota + 1
	//InstructionTypeAllow - Allow
	InstructionTypeAllow
	//InstructionTypeDeny - Deny
	InstructionTypeDeny
)

//Pointer to cpecific document section, explaining how to deal with it
type partShareInstructionType struct {
	b        *builder
	tagName  string
	linkType string
	href     string
	include  DocGenerationInstructionType
}

func (s *partShareInstructionType) builder() *builder {
	if s.b == nil {
		s.b = &builder{}
	}
	return s.b
}

func (s *partShareInstructionType) ToXML() (string, error) {
	if s.href == "" {
		return "", makeError(ErrEmptyAttribute, "Empty required attribute %s/href", s.tag())
	}
	if s.include == 0 {
		return "", makeError(ErrEmptyAttribute, "Empty required attribute %s/include", s.tag())
	}
	str, err := s.include.Value()
	if err != nil {
		return "", wrapError(err, ErrNestedEntity, "Can't make %s/include attribute", s.tag())
	}
	b := s.builder()
	b.Reset()
	b.makeTagAttr(s.tag(), "", map[string]string{"type": s.linkType, "href": s.href, "include": str}, false)
	return b.String(), nil
}

func (s *partShareInstructionType) tag() string {
	if s.tagName == "" {
		s.tagName = "part"
	}
	return s.tagName
}
