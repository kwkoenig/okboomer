// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func byteReader(blackbox chan byte) {
	// put stdin in raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)

	for {
		n, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n == 1 {
			blackbox <- b[0]
		} else {
			fmt.Println("bytes read <> 1 but no error was thrown.  go figure.  exiting.")
			os.Exit(1)
		}
	}
}
