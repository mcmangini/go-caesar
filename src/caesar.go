package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func usage() {
	fmt.Printf("Usage: %s [-bfh] [-s SIZE] <TEXT>\n", os.Args[0])
	fmt.Println("Encrypt/decrypt text using a Caesar cipher.\n" +
			"\n" +
			"Options:\n" +
			"  -b, --brute-force          decrypt using character frequency analysis;\n" +
			"                             overrides -s, --shift\n" +
			"  -f, --file                 read text from file (TEXT interpreted as file path)\n" +
			"  -h, --help                 display usage message\n" +
			"  -s, --shift <SIZE>         shift text by SIZE characters\n" +
			"\n" +
			"Exit status:\n" +
			" 0  if OK,\n" +
			" 1  if error.")
	os.Exit(0)
}

func caesarCipher(shift rune, in string) string {
	const span = 26
	shift %= span
	if shift < 0 {
		shift = span + shift
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
		out += string((c - start + shift) % span + start)
	}
	return out
}

func main() {
	var bruteForce bool
	var shift int
	var fp, in string
	args := os.Args[1:]

	// Parse args
	if len(args) > 6 {
		log.Fatal("Error: Too many arguments")
	}
	for i, v := range args {
		if v == "-b" || v == "--brute-force" {
			bruteForce = true
		} else if v == "-f" || v == "--file" {
			f, err := os.Open(args[len(args) - 1])
			defer f.Close()
			if err != nil {
				log.Fatal("Error: Invalid file")
			}
		} else if v == "-h" || v == "--help" {
			usage()
		} else if v == "-s" || v == "--shift" {
			if i < len(args) - 1 {
				n, err := strconv.Atoi(args[i + 1])
				if err != nil {
					log.Fatal("Error: Shift size must be an integer")
				}
				if n < math.MinInt32 || n > math.MaxInt32 {
					log.Fatal("Error: Invalid shift size")
				}
				shift = n
				i++
			}
		} else if i == len(args) - 1 && fp == "" {
			in = v
		}
	}
	fmt.Println(args)
	out := caesarCipher(rune(shift), in)
	fmt.Println(out)
}
