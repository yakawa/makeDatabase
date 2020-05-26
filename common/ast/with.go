package ast

type WithClause struct {
	IsRecursive bool
	CTE         []CTE
}

type CTE struct {
	TableName       string
	ColumnNames     []string
	SelectStatement *SelectStatement
}
