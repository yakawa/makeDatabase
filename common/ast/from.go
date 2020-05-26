package ast

type FromClause struct {
	ToS        *TableOrSubquery
	JoinClause *JoinClause
}

type TableOrSubquery struct {
	Schema          string
	TableName       string
	Alias           string
	TableOrSubquery *TableOrSubquery
	JoinClause      *JoinClause
	SelectStatement *SelectStatement
}

type JoinClause struct {
	LeftTableOrSubquery  *TableOrSubquery
	RightTableOrSubquery *TableOrSubquery
	Natural              bool
	Left                 bool
	Right                bool
	Cross                bool
	Expr                 *Expression
	ColumnNames          []string
}
