// Code generated by "stringer -type=DocGenerationInstructionType"; DO NOT EDIT.

package gofbwriter

import "strconv"

const _DocGenerationInstructionType_name = "InstructionTypeRequireInstructionTypeAllowInstructionTypeDeny"

var _DocGenerationInstructionType_index = [...]uint8{0, 22, 42, 61}

func (i DocGenerationInstructionType) String() string {
	if i < 0 || i >= DocGenerationInstructionType(len(_DocGenerationInstructionType_index)-1) {
		return "DocGenerationInstructionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DocGenerationInstructionType_name[_DocGenerationInstructionType_index[i]:_DocGenerationInstructionType_index[i+1]]
}
