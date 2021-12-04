package ast

// Every node in our AST has to implement Node interface, meaninig it has
// to provide a TokenLiteral() method that returns the literal value of
// the token it's associated with.
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// This Program node is going to be the root node of every AST our parser
// produces.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
