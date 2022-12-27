package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func raw() *term.State {
	// put stdin in raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return oldState
}
func cook(oldState *term.State) {
	term.Restore(int(os.Stdin.Fd()), oldState)
}
