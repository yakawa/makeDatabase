package compiler

import (
	"github.com/yakawa/makeDatabase/common/errors"
	"github.com/yakawa/makeDatabase/compiler/lexer"
	"github.com/yakawa/makeDatabase/compiler/parser"
	"github.com/yakawa/makeDatabase/logger"
	"github.com/yakawa/makeDatabase/tools/printer"
)

func Compile(sql string) (err error) {
	tokens, err := lexer.Tokenize(sql)
	if err != nil {
		return
	}
	a, err := parser.Parse(tokens)
	if err != nil {
		logger.Errorf("%+v", err)
		logger.Errorf("%s", err.(*errors.ErrParseInvalid).PrintStack(10))
	}
	x := printer.PrintSQL(a)
	logger.Infof("\n%s", x)
	return
}
