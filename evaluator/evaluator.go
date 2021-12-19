package evaluator

import (
	"github.com/petrostrak/the-monkey-programming-language/ast"
	"github.com/petrostrak/the-monkey-programming-language/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}