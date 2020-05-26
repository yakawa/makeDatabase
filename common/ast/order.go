package ast

type OrderClause struct {
	Expr        *Expression
	CollateName string
	Asc         bool
	Desc        bool
	NullsFirst  bool
	NullsLast   bool
}
