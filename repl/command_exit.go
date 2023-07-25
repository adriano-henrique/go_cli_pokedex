package repl

import "os"

func commandExit(mapConfig *MapConfig) error {
	os.Exit(1)
	return nil
}
