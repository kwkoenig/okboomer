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

const ctrlC byte = 3
const bell byte = 7
const backspace byte = 8
const carriageReturn byte = 13

func doyourthing(blackbox chan byte, flags int, whoknows []whadhesay) {
	fmt.Println("\nBegin typing.  Known translations will be preceded by =>.  Press Enter to clear, Ctrl-C to quit.")
	fmt.Println()
	var said, padded string
	var prevLen int
	var b byte

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
			w, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if tempLen > w-1 {
				writeabyte(bell)
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

		writeabyte(carriageReturn)
		_, err := os.Stdout.WriteString(padded)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for dif := len(padded) - len(said); dif > 0; dif-- {
			writeabyte(backspace)
		}
		prevLen = evalLen
	}

}

func writeabyte(b byte) {
	byteBuf := make([]byte, 1)
	byteBuf[0] = b
	n, err := os.Stdout.Write(byteBuf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if n != 1 {
		fmt.Println("bytes written <> 1 but without error")
		os.Exit(1)
	}
}
