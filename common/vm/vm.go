package vm

import "fmt"

type Operation interface {
	String() string
}

type GetTableOpe struct {
	Table string
}

func (o *GetTableOpe) String() string {
	return fmt.Sprintf("Get Table Operation Named By %s", o.Table)
}

type FilterOpe struct {
}
