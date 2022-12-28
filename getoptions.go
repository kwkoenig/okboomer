// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getoptions() int {
	var flags int
	var reply string
	var ok bool

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nGreetings, wise one.  What will we be translating today?  (m)illennial and/or (t)ext and/or (g)amer, or (a)ll three?: ")
	if reply, ok = readastring(reader); !ok {
		return 0
	}
	if !checksout(reply) {
		fmt.Print("\nLet's try that again.  Please enter any combination of the letters m, t, g, or a, per the above instructions: ")
		if reply, ok = readastring(reader); !ok {
			return 0
		}
		if !checksout(reply) {
			naptime()
			return 0
		}
	}

	if strings.Contains(reply, "a") {
		flags = all
	} else {
		if strings.Contains(reply, "m") {
			flags += millennial
		}
		if strings.Contains(reply, "t") {
			flags += text
		}
		if strings.Contains(reply, "g") {
			flags += gamer
		}
	}

	fmt.Print("\nShould the interpretations be (s)afe for work, or is (n)ot safe acceptable? ")
	if reply, ok = readastring(reader); !ok {
		return 0
	}
	if !sOrn(reply) {
		fmt.Print("\nLet's try that again.  Please enter either s or n, per the previous instructions: ")
		if reply, ok = readastring(reader); !ok {
			return 0
		}
		if !sOrn(reply) {
			naptime()
			return 0
		}
	}

	if strings.HasPrefix(reply, "n") {
		flags += nsfw
	}

	return flags
}

func readastring(reader *bufio.Reader) (string, bool) {
	red, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return red, false
	}
	return red, true
}

func checksout(input string) bool {
	if input == "" {
		return false
	}
	lower := strings.ToLower(input)
	return strings.Contains(lower, "m") ||
		strings.Contains(lower, "t") ||
		strings.Contains(lower, "g") ||
		strings.Contains(lower, "a")
}

func sOrn(reply string) bool {
	return strings.HasPrefix(reply, "s") || strings.HasPrefix(reply, "n")
}

func naptime() {
	fmt.Println("\nYou know, it may be time for a nap.  Good day to you.  Would you be so kind as to press any key?")
	fmt.Println()
	bb := make(chan byte)
	termstate := raw()
	go byteReader(bb)
	<-bb
	close(bb)
	cook(termstate)
}
