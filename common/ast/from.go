package ast

type FromClause struct {
	ToS []TableOrSubquery
}

type TableOrSubquery struct {
	Schema      string
	Database    string
	TableName   string
	Alias       string
	Subquery    *SelectStatement
	Natural     bool
	Left        bool
	Right       bool
	Inner       bool
	Cross       bool
	On          *Expression
	ColumnNames []string
}
