package ast

type SQL struct {
	SelectStatement *SelectStatement
}

type SelectStatement struct {
	WithClause   *WithClause
	SelectClause *SelectClause
	OrderClause  *OrderClause
	LimitClause  *LimitClause
}
