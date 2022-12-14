// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

import (
	"fmt"
	"os"
)

func doyourthing(flags int, whoknows []whadhesay) {

	b := make([]byte, 1)

	var err error
	var said string

	echo("", "")
	for {

		_, err = os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}

		hit := string(b[0])

		switch { // ya gotta love that
		case b[0] == ctrlC:
			return
		case hit == backspace:
			len := len(said)
			if len > 0 {
				said = said[:len-1]
			}
		case b[0] == enter:
			said = ""
		default:
			said += hit
		}

		means := translate(said, whoknows, flags)
		echo(said, means)
	}

}
