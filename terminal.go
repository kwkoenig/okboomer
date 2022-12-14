// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

import (
	"golang.org/x/term"
	"os"
)

func goraw() *term.State {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		os.Stdout.WriteString("Can't go raw.")
		os.Exit(1)
	}
	return oldState
}

func cookit(oldState *term.State) {
	term.Restore(int(os.Stdin.Fd()), oldState)
}
