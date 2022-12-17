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

const ctrlC = 3
const backspace = 8
const carriageReturn = 13
const becomes = "  =>  "

func doyourthing(flags int, whoknows []whadhesay) {
	fmt.Println("\nBegin typing.  Known translations will be preceded by =>.  Press Enter to clear, Ctrl-C to quit.")
	fmt.Println()
	// put stdin in raw mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	var said, means, output string
	prevLen := 0

	for {
		_, err = os.Stdin.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch {
		case b[0] == ctrlC:
			return
		case b[0] == carriageReturn:
			said = ""
		case b[0] == backspace:
			if len(said) > 0 {
				said = said[0 : len(said)-1]
			}
		default:
			said += string(b[0])
		}
		means = translate(said, whoknows, flags)
		if means != said {
			output = said + becomes + means
		} else {
			output = said
		}
		for dif := prevLen - len(output); dif > 0; dif-- {
			output += " "
		}
		b[0] = carriageReturn
		_, _ = os.Stdout.Write(b)
		_, _ = os.Stdout.WriteString(output)

		b[0] = backspace
		for dif := len(output) - len(said); dif > 0; dif-- {
			os.Stdout.Write(b)
		}
		prevLen = len(output)
	}

}
