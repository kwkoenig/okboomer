// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

// Bits in a byte.  Old school.
const millennial = 1
const text = 2
const gamer = 4
const all = 7
const nsfw = 8

type whadhesay struct {
	said  string
	means string
	flags int
}

func main() {

	var flags int
	// I just cant resist go's if with a statement
	// It's almost as good as C's non-zero == true
	if flags = getoptions(); flags == 0 {
		return
	}

	whoknows := damnkids()
	blackbox := make(chan byte)
	termstate := raw()

	go byteReader(blackbox)
	doyourthing(blackbox, flags, whoknows)
	close(blackbox)
	cook(termstate)
}
