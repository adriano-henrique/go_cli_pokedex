package repl

import "os"

func commandExit() error {
	os.Exit(1)
	return nil
}
