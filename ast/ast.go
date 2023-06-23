package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	// Used for debugging/help. This is simply to allow the Go compiler to guide us at times where
	// an expression may have been used in place of a statement or vice versa.
	statementNode()
}

// Expressions produce a value, statements do not.
// E.g. let x = 5; is a statement, but '5' is a value as the value produced itself is 5.
// Another example is add(1,1) is an expression, since it produces a value from the definition of 'add'.
type Expression interface {
	Node
	// Used for debugging/help. This is simply to allow the Go compiler to guide us at times where
	// an expression may have been used in place of a statement or vice versa.
	expressionNode()
}

// Program is always the root of our AST.
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

type LetStatement struct {
	Token token.Token // i.e. token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // i.e. token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token token.Token // i.e. token.RETURN
	Value string
}

func (ls *ReturnStatement) statementNode() {}

func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }
