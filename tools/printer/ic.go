package printer

import (
	"bytes"
	"fmt"

	"github.com/yakawa/makeDatabase/common/ic"
)

type icPrinter struct {
	tblList map[string]ic.TableInfo
}

func new(tblList map[string]ic.TableInfo) *icPrinter {
	return &icPrinter{
		tblList: tblList,
	}
}

func PrintIC(interCode []ic.Operation, tblList map[string]ic.TableInfo) string {
	p := new(tblList)
	return p.print(interCode)
}

func (p *icPrinter) print(interCode []ic.Operation) string {
	var out bytes.Buffer

	for _, ope := range interCode {
		switch ope.OpeCode() {
		case ic.GetTableOpeCode:
			str := p.printGetTableOpe(ope.(ic.GetTableOpe))
			out.WriteString(str)
		default:
			out.WriteString("Unknown Instruction Code\n")
			return out.String()
		}
	}
	return out.String()
}

func (p *icPrinter) printGetTableOpe(ope ic.GetTableOpe) string {
	var out bytes.Buffer
	t := p.tblList[ope.Table]
	out.WriteString("Get Table Operation:\n")
	if t.IsLocal {
		out.WriteString(fmt.Sprintf("\t Location: Schema - %s Database- %s Table - %s\n", t.Schema, t.Database, t.Table))
	} else {
		out.WriteString(fmt.Sprintf("\t Using Virtual Table"))
	}
	out.WriteString(fmt.Sprintf("\t ID: %s\n", ope.Table))

	return out.String()
}
