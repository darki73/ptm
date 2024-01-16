package debug

import (
	"github.com/Code-Hex/dd/p"
	"os"
)

// Dump dumps the arguments.
func Dump(args ...interface{}) {
	if _, err := p.P(args...); err != nil {
		panic(err)
	}
}

// DumpAndDie dumps the arguments and exits.
func DumpAndDie(args ...interface{}) {
	Dump(args...)
	os.Exit(0)
}
