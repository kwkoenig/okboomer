// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

import "os"

func echo(said, means string) {
	fp := os.Stdout
	fp.WriteString(clear)
	fp.WriteString("Press enter to clear, or ctrl-c to quit.\n\n")
	fp.WriteString("\n" + means)
	fp.WriteString(prevline + cr + said)
}
