package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func usage() {
	fmt.Printf("Usage: %s [-bfh] [-s <SIZE>] <TEXT>\n", os.Args[0])
	fmt.Println("Encrypt/decrypt text using a Caesar cipher.\n" +
			"\n" +
			"Options:\n" +
			"  -b, --brute-force          attempt to decrypt using English-language\n" + 
			"                             character frequency analysis;\n" +
			"                             unreliable for small inputs;\n" + 
			"                             overrides -s, --shift\n" +
			"  -f, --file                 read text from file (interpret TEXT as file path)\n" +
			"  -h, --help                 display usage message\n" +
			"  -s, --shift <SIZE>         shift text by SIZE characters\n" +
			"\n" +
			"Exit status:\n" +
			" 0  if OK,\n" +
			" 1  if error.")
	os.Exit(0)
}

// Shift each character in "in" by "shift" characters and return result
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

// Assign score to "text" based on character frequency and return score;
// the higher the score, the more likely the text is English
func scoreEnglishText(text string) float64 {
	freq := map[rune]float64{
		'a': 0.06078062040132936,
		'b': 0.010608146612806431,
		'c': 0.017582744552198538,
		'd': 0.035354605100584496,
		'e': 0.09437770751067875,
		'f': 0.01708473909557166,
		'g': 0.015823965787655516,
		'h': 0.04920041756811958,
		'i': 0.05167909789148272,
		'j': 0.0009001921418521265,
		'k': 0.006039104144918327,
		'l': 0.02778113983993222,
		'm': 0.019204099026178697,
		'n': 0.05344165897594949,
		'o': 0.05868395439026481,
		'p': 0.01259638611939019,
		'q': 0.0008396750230721515,
		'r': 0.04698649963941883,
		's': 0.04737355704494909,
		't': 0.06816623043910212,
		'u': 0.02110282362790041,
		'v': 0.006556021201163946,
		'w': 0.01782733457393427,
		'x': 0.0009115391016233717,
		'y': 0.015380173583269034,
		'z': 0.0002698054878940547,
	}
	var score float64
	for _, c := range strings.ToLower(text) {
		score += freq[c]
	}
	return score / float64(len(text))
}

// Attempt to decode "in" by scoring all 26 possible shifts and returning
// the one with the highest score;
// unreliable for small "in" strings
func bruteForceDecode(in string) string {
	var bestResult string
	var maxScore float64

	for i := range 26 {
		t := caesarCipher(rune(i), in)
		score := scoreEnglishText(t)
		if score > maxScore {
			bestResult = t
			maxScore = score
		}
	}
	return bestResult
}

// Set flags according to provided arguments
func parseArgs(bf, f *bool, s *rune, in *string, args []string) {
	var skip bool
	for i, v := range args {
		if skip {
			skip = false
			continue
		}
		if v == "-b" || v == "--brute-force" {
			*bf = true
		} else if v == "-f" || v == "--file" {
			*f = true
		} else if v == "-s" || v == "--shift" {
			if i < len(args) - 1 {
				n, err := strconv.Atoi(args[i + 1])
				if err != nil {
					log.Fatal("Shift size must be an integer")
				}
				// Shift must be within bounds of a rune
				if n < math.MinInt32 || n > math.MaxInt32 {
					log.Fatal("Invalid shift size")
				}
				*s = rune(n)
				skip = true
			} else {
				log.Fatal("No shift size provided")
			}
		} else if i == len(args) - 1 {
			*in = v
		} else {
			log.Fatalf("Unknown flag '%s'\n", v)
		}
	}
}

func main() {
	var bruteForce, fromFile bool
	var shift rune
	var in string
	args := os.Args[1:]

	log.SetFlags(0)
	log.SetPrefix("caesar: ")

	// Prioritize help argument over all others
	for _, v := range args {
		if v == "-h" || v == "--help" {
			usage()
		}
	}
	parseArgs(&bruteForce, &fromFile, &shift, &in, args)

	// Handle flags
	if in == "" {
		// If no text provided, read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		var lines string

		// Read until Ctl-D is pressed
		for scanner.Scan() {
			lines += scanner.Text() + "\n"
		}
		// Remove trailing newline
		in = lines[:len(lines)-1]
	}
	if fromFile {
		data, err := os.ReadFile(in)
		if err != nil {
			log.Fatalf("Could not read file '%s'", in)
		}
		// Remove trailing newline
		in = string(data[:len(data)-1])
	}

	if bruteForce {
		fmt.Println(bruteForceDecode(in))
	} else {
		fmt.Println(caesarCipher(shift, in))
	}
}
