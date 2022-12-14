// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

import "strings"

func translate(said string, whoknows []whadhesay, flags int) string {
	atwork := !isnsfw(flags)
	languages := flags & all
	means := strings.ToLower(said)
	for i, len := 0, len(whoknows); i < len; i++ {
		if strings.Contains(means, whoknows[i].said) {
			if whoknows[i].flags&languages > 0 {
				pottymouth := isnsfw(whoknows[i].flags)
				if !pottymouth || !atwork {
					means = strings.ReplaceAll(means, whoknows[i].said, whoknows[i].means)
				}
			}
		}
	}
	return means
}

func isnsfw(flags int) bool {
	return flags&nsfw > 0
}
