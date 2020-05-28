package ast

import "github.com/yakawa/makeDatabase/common/token"

type Expression struct {
	Literal    *Literal
	ColumnName *ColumnName
	PrefixOpe  *UnaryOpe
	BinaryOpe  *BinaryOpe
	PostfixOpe *UnaryOpe
	Function   *Function
	Expr       []Expression
	Cast       *CastOpe
	Collate    *CollateOpe
	String     *StringOpe
	Null       *NullOpe
	Not        *NotOpe
	Between    *BetweenOpe
	In         *InOpe
	Exists     *ExistsOpe
	Case       *CaseOpe
}

func (n *Expression) expressionNode() {
}

type Literal struct {
	Numeric          *Numeric
	String           string
	IsString         bool
	Null             bool
	True             bool
	False            bool
	CurrentTime      bool
	CurrentDate      bool
	CurrentTimestamp bool
}

type Numeric struct {
	SignedInt   bool
	UnsignedInt bool
	Float       bool
	SI          int
	UI          uint
	FL          float64
}

func (n *Numeric) expressionNode() {
}

type ColumnName struct {
	Schema string
	Table  string
	Column string
}

type UnaryOpe struct {
	Operator token.Type
	Expr     *Expression
}

type BinaryOpe struct {
	Operator token.Type
	Expr1    *Expression
	Expr2    *Expression
}

type Function struct {
	Name       string
	Args       []Expression
	Distinct   bool
	Asterisk   bool
	FilterExpr *Expression
	OverClause *OverClause
}

type OverClause struct {
	WindowName     string
	BaseWindowName string
	PartitionExpr  []Expression
	OrderBy        *OrderClause
	FrameSpec      *FrameSpecification
}

type CastOpe struct {
	Expr     *Expression
	TypeName string
	N1       int
	N2       int
}

type CollateOpe struct {
	Name string
	Expr *Expression
}

type StringOpe struct {
	Not        bool
	Like       bool
	Glob       bool
	Regexp     bool
	Match      bool
	Expr1      *Expression
	Expr2      *Expression
	EscapeExpr *Expression
}

type NullOpe struct {
	Expr    *Expression
	IsNull  bool
	NotNull bool
}

type NotOpe struct {
	Expr1 *Expression
	Expr2 *Expression
	Not   bool
}

type BetweenOpe struct {
	Expr1 *Expression
	Expr2 *Expression
	Expr3 *Expression
	Not   bool
}

type InOpe struct {
	Expr            *Expression
	Not             bool
	SelectStatement *SelectStatement
	InExpr          []Expression
	Schema          string
	Table           string
}

type ExistsOpe struct {
	Not             bool
	SelectStatement *SelectStatement
}

type CaseOpe struct {
	CaseExpr *Expression
	WhenThen []WhenThen
	ElseExpr *Expression
}

type WhenThen struct {
	WhenExpr *Expression
	ThenExpr *Expression
}
