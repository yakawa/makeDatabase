package ast

type SelectClause struct {
	IsDistinct        bool
	IsAll             bool
	ResultColumns     []ResultColumn
	FromClause        *FromClause
	WhereClause       *WhereClause
	GroupByExpression *GroupByExpression
	WindowExpression  *WindowExpression
}

type ResultColumn struct {
	Expr       *Expression
	Alias      string
	Asterisk   bool
	SchemaName string
	TableName  string
}

type GroupByExpression struct {
	GroupingExpr []Expression
	HavingExpr   *Expression
}

type WindowExpression struct {
	WindowDefn []WindowDefinition
}

type WindowDefinition struct {
	WindowName    string
	PartitionExpr []Expression
	GroupExpr     []Expression
	Frame         *FrameSpecification
}

type FrameSpecification struct {
	Range  bool
	Rows   bool
	Groups bool

	Between             bool
	UnboundedPreceding1 bool
	ExprPreceding1      *Expression
	CurrentRow1         bool
	ExprFollowing1      *Expression
	ExprPreceding2      *Expression
	CurrentRow2         bool
	ExprFollowing2      *Expression
	UnboundedFollowing2 bool

	UnboundedPreceding bool
	ExprPreceding      *Expression
	CurrentRow         bool

	ExcludeNoOthers   bool
	ExcludeCurrentRow bool
	ExcludeGroup      bool
	ExcludeTies       bool
}

type CompoundOperator struct {
	Union           bool
	UnionAll        bool
	Intersect       bool
	Except          bool
	SelectStatement *SelectStatement
}
