package ast

type LimitClause struct {
	LimitExpr        *Expression
	OffsetExpression *Expression
}
