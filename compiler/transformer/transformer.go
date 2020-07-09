package transformer

import (
	"errors"

	"github.com/google/uuid"

	"github.com/yakawa/makeDatabase/common/ast"
	"github.com/yakawa/makeDatabase/common/vm"
	"github.com/yakawa/makeDatabase/logger"
)

type TableInfo struct {
	IsLocal  bool
	Schema   string
	Database string
	Table    string
}

type transformer struct {
	a     *ast.SQL
	tl    map[string]TableInfo
	revTl map[TableInfo]string
	atl   map[string]string
}

func Transform(a *ast.SQL) ([]vm.Operation, map[string]TableInfo, error) {
	t := new(a)
	ops, err := t.transform()
	return ops, t.tl, err
}

func new(a *ast.SQL) *transformer {
	return &transformer{
		a:     a,
		tl:    make(map[string]TableInfo),
		revTl: make(map[TableInfo]string),
		atl:   make(map[string]string),
	}
}

func (t *transformer) transform() ([]vm.Operation, error) {
	if t.a.SelectStatement.WithClause != nil {
		logger.Errorf("Not Impliment")
		return nil, errors.New("Not Impliment")
	}

	if t.a.SelectStatement.SelectClause != nil {
		t.transformSelectClause(t.a.SelectStatement.SelectClause)
	}

	return nil, nil
}

func (t *transformer) transformSelectClause(s *ast.SelectClause) ([]vm.Operation, error) {
	if s.FromClause != nil {
		for _, tos := range s.FromClause.ToS {
			t.transformToS(tos)
		}
	}
	return nil, nil
}

func (t *transformer) transformToS(tos ast.TableOrSubquery) ([]vm.Operation, error) {
	ope := []vm.Operation{}
	tbl := TableInfo{}
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

		tuuid, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		tid := tuuid.String()
		getTableOpe := vm.GetTableOpe{
			Table: tid,
		}

		t.tl[tid] = tbl
		t.revTl[tbl] = tid

		ope = append(ope, &getTableOpe)
	}
	return ope, nil
}
