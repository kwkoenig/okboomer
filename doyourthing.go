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
const bell = 7
const backspace = 8
const carriageReturn = 13

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
	var said, padded string
	var prevLen int

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
				said = said[:len(said)-1]
			}
		default:
			_, tempLen := translate(said+string(b[0]), whoknows, flags)
			w, _, _ := term.GetSize(int(os.Stdout.Fd()))
			if tempLen > w-1 {
				b[0] = bell
				_, _ = os.Stdout.Write(b)
			} else {
				said += string(b[0])
			}
		}
		eval, evalLen := translate(said, whoknows, flags)
		padded = eval
		if evalLen < prevLen {
			for i := evalLen; i < prevLen; i++ {
				padded += " "
			}
		}

		b[0] = carriageReturn
		_, _ = os.Stdout.Write(b)
		_, _ = os.Stdout.WriteString(padded)

		b[0] = backspace
		for dif := len(padded) - len(said); dif > 0; dif-- {
			os.Stdout.Write(b)
		}
		prevLen = evalLen
	}

}
