package compiler

import (
	"testing"

	"github.com/petrostrak/the-monkey-programming-language/ast"
	"github.com/petrostrak/the-monkey-programming-language/code"
	"github.com/petrostrak/the-monkey-programming-language/lexer"
	"github.com/petrostrak/the-monkey-programming-language/parser"
)

type compilerTesCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	testCases := []compilerTesCase{
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
			},
		},
	}

	runCompilerTests(t, testCases)
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func runCompilerTests(t *testing.T, testCases []compilerTesCase) {
	t.Helper()

	for _, tt := range testCases {
		program := parse(tt.input)

		compiler := New()
		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		err = testConstants(t, tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}
	}
}
