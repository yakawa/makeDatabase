package repl

import (
	"io"
	"strings"

	"github.com/chzyer/readline"

	"github.com/yakawa/makeDatabase/compiler"
	"github.com/yakawa/makeDatabase/logger"
)

const (
	PROMPT1 = " > "
	PROMPT2 = " | "
)

func Start(i io.Reader, w io.Writer, we io.Writer) {
	ioReader, ioWriter := readline.NewFillableStdin(i)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 PROMPT1,
		HistoryFile:            ".makeDatabase.readline",
		DisableAutoSaveHistory: true,
		Stdin:                  ioReader,
		StdinWriter:            ioWriter,
		Stdout:                 w,
		Stderr:                 we,
	})
	if err != nil {
		logger.Panicf("%s", err)
	}

	for {
		sql := ""
		for {
			line, err := rl.Readline()
			if err != nil {
				if err == io.EOF {
					return
				}

				sql = ""
				rl.SetPrompt(PROMPT1)
				continue
			}
			line = strings.TrimSpace(line)
			if len(line) == 0 {
				rl.SetPrompt(PROMPT1)
				break
			}
			rl.SetPrompt(PROMPT2)
			sql += line
			sql += "\n"
		}
		logger.Infof("SQL: %s", sql)
		compiler.Compile(sql)
	}
}
