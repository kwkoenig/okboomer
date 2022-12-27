// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

import (
	"fmt"
	"os"
)

func byteReader(blackbox chan byte) {
	// terminal MUST be in raw mode
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
