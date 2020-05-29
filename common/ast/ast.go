package ast

type SQL struct {
	SelectStatement *SelectStatement
}

type SelectStatement struct {
	WithClause      *WithClause
	SelectClause    *SelectClause
	ValuesClause    *ValuesClause
	CompoundOpeator *CompoundOperator
	OrderClause     *OrderClause
	LimitClause     *LimitClause
}
