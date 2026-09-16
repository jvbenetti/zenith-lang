package ast

// Node is the base interface for all nodes in the AST. Every node in our AST must have a TokenLiteral() method.
// This is only for debugging purposes, so we can print which token generated this node.
type Node interface {
	TokenLiteral() string
}

// Statement is a node that does not produce a value (e.g., let x = 5;)
type Statement interface {
	Node
	statementNode() // Dummy method to distinguish Statement from Expression
}

// Expression is a node that produces a value (e.g., 5 + 5)
type Expression interface {
	Node
	expressionNode() // Dummy method to distinguish Expression from Statement
}

// Program is the ROOT node of our entire tree.
// A program in Zenith is simply a list of Statements (declarations) being executed one after another.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
