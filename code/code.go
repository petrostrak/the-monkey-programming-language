package code

import "fmt"

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
)

// The Definition of an Opcode has two fields. Name helps
// to make an Opcode readable and OperandWidths contains
// the number of bytes each operand takes up.
type Definition struct {
	Name          string
	OperandWidths []int
}

var (
	definitions = map[Opcode]*Definition{
		OpConstant: {"OpConstant", []int{2}},
	}
)

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}
