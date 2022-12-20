// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

import "strings"

const becomes = "  =>  "

func translate(said string, whoknows []whadhesay, flags int) (string, int) {
	notatwork := isnsfw(flags)
	languages := flags & all
	means := strings.ToLower(said)
	for i, len := 0, len(whoknows); i < len; i++ {
		if strings.Contains(means, whoknows[i].said) {
			if whoknows[i].flags&languages > 0 {
				pottymouth := isnsfw(whoknows[i].flags)
				if notatwork || !pottymouth {
					means = strings.ReplaceAll(means, whoknows[i].said, whoknows[i].means)
				}
			}
		}
	}
	if means != said {
		r := said + becomes + means
		return r, len(r)
	}
	return said, len(said)
}

func isnsfw(flags int) bool {
	return flags&nsfw > 0
}
