package errors

import (
	"bytes"
	"fmt"
	"runtime"

	"github.com/yakawa/makeDatabase/common/token"
)

type StackTrace struct {
	File string
	Line int
	Func string
}

type ErrParseInvalid struct {
	Token token.Token
	Trace []StackTrace
}

func (e *ErrParseInvalid) Error() string {
	return fmt.Sprintf("ParseError: Token %s (%s) is invalid", e.Token.Type.String(), e.Token.Literal)
}

func (e *ErrParseInvalid) PrintStack(n int) string {
	var out bytes.Buffer

	if n > len(e.Trace) {
		n = len(e.Trace)
	}
	for i := 0; i < n; i++ {
		out.WriteString(fmt.Sprintf("Func: %s (%s:%d)\n", e.Trace[i].Func, e.Trace[i].File, e.Trace[i].Line))
	}

	return out.String()
}

func NewErrParseInvalidToken(t token.Token) *ErrParseInvalid {
	e := &ErrParseInvalid{
		Token: t,
	}

	for i := 1; ; i++ {
		pt, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pt).Name()
		e.Trace = append(e.Trace, StackTrace{File: file, Line: line, Func: fn})
	}
	return e
}
