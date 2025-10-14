package main

import (
	"fmt"
	"os"
)

func caesarCipher(shift rune, in string) string {
	shift %= 26
	if shift < 0 {
		shift = 26 + shift
	}
	ucStart := 'A'
	ucEnd := 'Z'
	lcStart := 'a'
	lcEnd := 'z'
	out := ""
	for _, c := range in {
		var start rune
		if c >= ucStart && c <= ucEnd {
			// A - Z
			start = ucStart
		} else if c >= lcStart && c <= lcEnd {
			// a - z
			start = lcStart
		} else {
			out += string(c)
			continue
		}
		out += string((c - start + shift) % 26 + start)
	}
	return out
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	out := caesarCipher(-72, "This is a test")
	fmt.Println(out)
}
