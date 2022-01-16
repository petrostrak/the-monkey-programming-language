package vm

import (
	"fmt"

	"github.com/petrostrak/the-monkey-programming-language/ast"
	"github.com/petrostrak/the-monkey-programming-language/lexer"
	"github.com/petrostrak/the-monkey-programming-language/object"
	"github.com/petrostrak/the-monkey-programming-language/parser"
)

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Object)
	if !ok {
		return fmt.Errorf("object is not Integer; got %T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value; got %d, want %d", result.Value, expected)
	}

	return nil
}
