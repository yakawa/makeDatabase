package compiler

import (
	"github.com/yakawa/makeDatabase/compiler/lexer"
	"github.com/yakawa/makeDatabase/compiler/parser"
)

func Compile(sql string) (err error) {
	tokens, err := lexer.Tokenize(sql)
	if err != nil {
		return
	}
	parser.Parse(tokens)
	return
}
