package ast

import "monkey/token"

//every node in our AST implement Node interface
type Node interface {
	//TokenLiteral returns the literal value of the token it's associated with
	//will be used for debugging
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	expressionNode()
}

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

//Name holds the identifier of the binding
//Value holds the expression that produces value
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) StatementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
