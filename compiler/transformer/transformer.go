package transformer

import (
	"errors"

	"github.com/google/uuid"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/ic"
	"github.com/yakawa/makeDatabase/logger"
)

type transformer struct {
	a     *ast.SQL
	opes  []ic.Operation
	tl    map[string]ic.TableInfo
	revTl map[ic.TableInfo]string
	atl   map[string]string
}

func Transform(a *ast.SQL) ([]ic.Operation, map[string]ic.TableInfo, error) {
	t := new(a)
	err := t.transform()
	return t.opes, t.tl, err
}

func new(a *ast.SQL) *transformer {
	return &transformer{
		a:     a,
		opes:  []ic.Operation{},
		tl:    make(map[string]ic.TableInfo),
		revTl: make(map[ic.TableInfo]string),
		atl:   make(map[string]string),
	}
}

func (t *transformer) transform() error {
	if t.a.SelectStatement.WithClause != nil {
		logger.Errorf("Not Impliment")
		return errors.New("Not Impliment")
	}

	if t.a.SelectStatement.SelectClause != nil {
		t.a.SelectStatement.SelectClause.FromClause != nil {
			err := t.transformFromClause(t.a.SelectStatement.SelectClause.FromClause)
			if err != nil {
				return err
			}
		}
		t.a.SelectStatement.SelectClause.WhereClause != nil {
			err := t.transformWhereClause(t.a.SelectStatement.SelectClause.WhereClause)
		}
	}

	return nil
}

func (t *transformer) transformWhereClayse(w *ast.WhereClause) error {
	return nil
}

func (t *transformer) transformFromClause(f *ast.FromClause) error {
	if f != nil {
		for _, tos := range f.ToS {
			err := t.transformToS(tos)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *transformer) transformToS(tos ast.TableOrSubquery) error {
	ope := ic.GetTableOpe{}
	tbl := ic.TableInfo{}
	if tos.TableName != "" {
		if tos.Schema == "" {
			tbl.Schema = "."
		} else {
			tbl.Schema = tos.Schema
		}
		if tos.Database == "" {
			tbl.Database = "."
		} else {
			tbl.Database = tos.Database
		}
		tbl.Table = tos.TableName
		tbl.IsLocal = true

		tuuid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		tid := tuuid.String()

		t.tl[tid] = tbl
		t.revTl[tbl] = tid

		if tos.Alias != "" {
			t.atl[tos.Alias] = tid
		}
		ope.Table = tid
	}

	t.opes = append(t.opes, ope)

	return nil
}
