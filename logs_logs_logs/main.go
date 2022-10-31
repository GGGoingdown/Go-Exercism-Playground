package main

import (
	"fmt"
	"strconv"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	runeLog := []rune(log)
	for _, s := range runeLog {
		switch s {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runeLog := []rune(log)
	for i, v := range runeLog {
		if v == oldRune {
			runeLog[i] = newRune
		}
	}
	return string(runeLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runeLog := []rune(log)
	if len(runeLog) > limit {
		return false
	}
	return true
}

func main() {
	// a := Application("❗ recommended search product 🔍")
	// fmt.Println(a)
	log := "please replace '👎' with '👍'"
	r := Replace(log, '👎', '👍')
	fmt.Println(r)
}
