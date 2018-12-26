package gofbwriter

import (
	"strings"
)

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
	tagName  string
	linkType string
	href     string
	include  DocGenerationInstructionType
}

func (s *partShareInstructionType) ToXML() (string, error) {
	if s.tagName == "" {
		s.tagName = "part"
	}
	if s.href == "" {
		return "", makeError(ErrEmptyAttribute, "Empty required attribute %s/href", s.tagName)
	}
	if s.include == 0 {
		return "", makeError(ErrEmptyAttribute, "Empty required attribute %s/include", s.tagName)
	}
	var b strings.Builder
	b.WriteString("<")
	b.WriteString(s.tagName)
	if s.linkType != "" {
		b.WriteString(" ")
		b.WriteString(makeAttribute("type", s.linkType))
	}
	b.WriteString(" ")
	b.WriteString(makeAttribute("href", s.href))
	b.WriteString(" ")
	str, err := s.include.Value()
	if err != nil {
		return "", wrapError(err, ErrNestedEntity, "Can't make %s/include attribute", s.tagName)
	}
	b.WriteString(makeAttribute("include", str))
	b.WriteString(">\n")
	return b.String(), nil
}

func (s *partShareInstructionType) tag() string {
	if s.tagName == "" {
		s.tagName = "part"
	}
	return s.tagName
}
