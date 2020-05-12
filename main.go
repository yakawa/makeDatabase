package main

import (
	"os"

	"github.com/yakawa/makeDatabase/logger"
	"github.com/yakawa/makeDatabase/repl"
)

func main() {
	logger.SetLevel("TRACE")
	logger.Infof("makeDatabase System")
	repl.Start(os.Stdin, os.Stdout, os.Stderr)
}
