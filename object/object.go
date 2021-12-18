package object

type ObjectType string

// Every value will be wrapped inside a struct, which
// fulfills this Object interface.
type Object interface {
	Type() ObjectType
	Inspect() string
}
