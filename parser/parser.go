package parser

import (
	"fmt"

	"github.com/petrostrak/the-monkey-programming-language/ast"
	"github.com/petrostrak/the-monkey-programming-language/lexer"
	"github.com/petrostrak/the-monkey-programming-language/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string

	// In order for our parser to get the correct prefixParserFn or infixParseFn
	// for the current token type, we add two maps to the Parser structure.
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken advances both curToken and peekToken.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectedPeek is one of the assertion functions. The primary purpose
// of it is to enforce the correctness of the order of tokens by checking
// the type of the next token.
func (p *Parser) expectedPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) parseLetStatement() *ast.LetStatement {

	// We construct a node with the token it's currently sitting
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	// TODO: we skip the expression until we encounter a semicolon

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: We are skipping the expression until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseStatement 's job is to parse a statement.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {

	// First thing, we construct the root node of the AST.
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Then we iterate over every token in the input until it
	// encounters an token.EOF.
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// Implementing the PRATT Parser.

type (
	prefixParseFn func() ast.Expression

	// the ast.Expression argument in this func is the "left side" of
	// the infix operator that's being parse.
	infixParseFn func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
