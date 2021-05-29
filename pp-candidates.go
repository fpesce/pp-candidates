package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var avoidDoublon bool
var wordlist string
var separator string
var endOfPhrase string
var minWords int
var maxWords int
var maxLength int

func init() {
	flag.BoolVar(&avoidDoublon, "a", false, "avoid phrase with twice the same word")
	flag.StringVar(&wordlist, "w", "", "wordlist to generate pass from")
	flag.StringVar(&separator, "s", "", "separator between words")
	flag.StringVar(&endOfPhrase, "e", "", "add a last string/char to every phrases")
	flag.IntVar(&minWords, "m", 4, "minimum number of words in the pass phrase")
	flag.IntVar(&maxWords, "n", 4, "maximum number of words in the pass phrase")
	flag.IntVar(&maxLength, "l", 64, "maximum length of the generated passphrase")
}

func recurse(prefix string, words []string, wordPresence []bool, depth int) {
	if depth > 0 && len(prefix) <= maxLength {
		for i, w := range words {
			if avoidDoublon {
				if wordPresence[i] {
					continue
				}
				wordPresence[i] = true
			}
			if depth == 1 {
				output := prefix + separator + w + endOfPhrase
				if len(output) <= maxLength {
					fmt.Println(output)
				}
				if avoidDoublon {
					wordPresence[i] = false
				}
			} else {
				var currentPrefix string
				if prefix != "" {
					currentPrefix = prefix + separator + w
				} else {
					currentPrefix = w
				}
				recurse(currentPrefix, words, wordPresence, depth - 1)
				if avoidDoublon {
					wordPresence[i] = false
				}
			}
		}
	}
}

func main() {
	flag.Parse()
	wl, err := os.Open(wordlist)
	if err != nil {
		log.Fatalf("can't open wordlist(%s):%s\n", wordlist, err)
	}
	defer wl.Close()
	var words []string
	scanWl := bufio.NewScanner(wl)
	for scanWl.Scan() {
		words = append(words, scanWl.Text())
	}
	var wordPresence []bool
	if avoidDoublon {
		wordPresence = make([]bool, len(words))
	}
	for i := minWords; i <= maxWords; i++ {
		recurse("", words, wordPresence, i)
	}
}
