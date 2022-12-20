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

func doyourthing(blackbox chan byte, flags int, whoknows []whadhesay) {
	fmt.Println("\nBegin typing.  Known translations will be preceded by =>.  Press Enter to clear, Ctrl-C to quit.")
	fmt.Println()
	var said, padded string
	var prevLen int
	var b byte

	byteArr := make([]byte, 1)

	for {
		b = <-blackbox
		switch {
		case b == ctrlC:
			return
		case b == carriageReturn:
			said = ""
		case b == backspace:
			if len(said) > 0 {
				said = said[:len(said)-1]
			}
		default:
			_, tempLen := translate(said+string(b), whoknows, flags)
			w, _, _ := term.GetSize(int(os.Stdout.Fd()))
			if tempLen > w-1 {
				byteArr[0] = bell
				_, _ = os.Stdout.Write(byteArr)
			} else {
				said += string(b)
			}
		}
		eval, evalLen := translate(said, whoknows, flags)
		padded = eval
		if evalLen < prevLen {
			for i := evalLen; i < prevLen; i++ {
				padded += " "
			}
		}

		byteArr[0] = carriageReturn
		_, _ = os.Stdout.Write(byteArr)
		_, _ = os.Stdout.WriteString(padded)

		byteArr[0] = backspace
		for dif := len(padded) - len(said); dif > 0; dif-- {
			os.Stdout.Write(byteArr)
		}
		prevLen = evalLen
	}

}
