package object

import "fmt"

const (
	INTEGER_OBJ = "INTEGER"
)

type ObjectType string

// Every value will be wrapped inside a struct, which
// fulfills this Object interface.
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Whenever we encounter an integer literal in the source code, we
// first turn it into an ast.IntegerLiteral and then, when evaluating
// that AST node, we turn it into an object.Integer, saving the value
// inside the Integer struct and passing around a reference to this
// struct.
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
