package vm

type OpeCode int
type JoinMethod int

const (
	GetTableOpeCode OpeCode = iota + 1
)

const (
	InnerJoin JoinMethod = iota + 1
	LeftJoin
	RightJoin
	FullJoin
)

type Operation interface {
	OpeCode() OpeCode
}

type GetTableOpe struct {
	Table string
}

func (o *GetTableOpe) OpeCode() OpeCode {
	return GetTableOpeCode
}

type JoinTableOpe struct {
	LeftTable  string
	RightTable string
	Method     JoinMethod
}
