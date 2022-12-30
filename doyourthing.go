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
	var said, eval, padded string
	var evalLen, prevLen int
	var b byte

	for ; ; prevLen = evalLen {
		b = <-blackbox
		switch b {
		case ctrlC:
			return
		case carriageReturn:
			said, eval, evalLen = "", "", 0
		case backspace:
			if len(said) > 0 {
				said = said[:len(said)-1]
				eval, evalLen = translate(said, whoknows, flags)
			}
		default:
			temp, tempLen := translate(said+string(b), whoknows, flags)
			// MUST check terminal width each time since user can resize.
			w, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				fmt.Println(err)
				return
			}
			if tempLen > w-1 {
				if !writeabyte(bell) {
					return
				}
			} else {
				said += string(b)
				eval, evalLen = temp, tempLen
			}
		}
		padded = eval
		if evalLen < prevLen {
			for i := evalLen; i < prevLen; i++ {
				padded += " "
			}
		}

		if !writeabyte(carriageReturn) {
			return
		}
		_, err := os.Stdout.WriteString(padded)
		if err != nil {
			fmt.Println(err)
			return
		}
		if diff := len(padded) - len(said); diff > 0 {
			if !backup(diff) {
				return
			}
		}
	}

}

func writeabyte(b byte) bool {
	byteBuf := make([]byte, 1)
	byteBuf[0] = b
	return writeit(byteBuf)
}

func backup(howmany int) bool {
	buf := make([]byte, howmany)
	for i := 0; i < howmany; i++ {
		buf[i] = backspace
	}
	return writeit(buf)
}

func writeit(b []byte) bool {
	n, err := os.Stdout.Write(b)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if n != len(b) {
		fmt.Println("bytes written != buffer length but no error was thrown.  go figure.  exiting.")
		return false
	}
	return true
}
