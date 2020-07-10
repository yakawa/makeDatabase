package vm

import "fmt"

type OpeCode int

const (
	GetTableOpeCode OpeCode = iota + 1
)

type Operation interface {
	OpeCode() OpeCode
	String() string
}

type GetTableOpe struct {
	Table string
}

func (o *GetTableOpe) String() string {
	return fmt.Sprintf("Get Table Operation Named By %s", o.Table)
}

func (o *GetTableOpe) OpeCode() OpeCode {
	return GetTableOpeCode
}

type FilterOpe struct {
}
