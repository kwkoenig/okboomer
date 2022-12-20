// Copyright 2022 by Keith W. Koenig
// No rights reserved.  Have at it.
// A little go program to help old-
// timers bridge the generation gap.

package main

func damnkids() []whadhesay {
	// Processes in order.
	// If one 'said' is a subset of another, the superset must proceed the subset.
	results := []whadhesay{
		{"bc", "because", text},
		{"btw", "by the way", text},
		{"cya", "see you", text},
		{"dm", "direct message", text},
		{"ftw", "for the win", text},
		{"fwiw", "for what it's worth", text},
		{"ikr", "i know, right?", text},
		{"imho", "in my humble opinion", text},
		{"imo", "in my opinion", text},
		{"irl", "in real life", text},
		{"jk", "just kidding", text},
		{"lmao", "laughing my ass off", text + nsfw},
		{"lmao", "laughing my butt off", text},
		{"lmk", "let me know", text},
		{"lol", "laughing out loud", text},
		{"nbd", "no big deal", text},
		{"np", "no problem", text},
		{"nsfw", "not safe for work", text},
		{"nvm", "never mind", text},
		{"omg", "oh my gosh", text},
		{"omg", "oh my god", text + nsfw},
		{"omw", "on my way", text},
		{"rofl", "rolling on the floor laughing", text},
		{"thx", "thanks", text},
		{"tmi", "too much information", text},
		{"ttyl", "talk to you later", text},
		{"yolo", "you only live once", text},
		{"wtf", "what the fuck", text + nsfw},
		{"wtf", "what the flock", text},
		{"a thing", "something people know about", millennial},
		{"i can't even", "i'm overwhelmed", millennial},
		{"i cant even", "i'm overwhelmed", millennial},
		{"i literally can't even", "i'm figuatively overwhelmed", millennial},
		{"i literally cant even", "i'm figuratively overwhelmed", millennial},
		{"literally", "figuratively", millennial},
		{"ok boomer", "sure old man", millennial},
		{"boomer", "old man", millennial},
		{"feel like", "think", millennial},
		{"go touch grass", "go outside", gamer},
		{"touch grass", "go outside", gamer},
		{"try hard", "loser that plays too much", gamer},
		{"tryhard", "loser that plays too much", gamer},
		{"uninstall", "quit playing the game", gamer},
	}
	return results
}
