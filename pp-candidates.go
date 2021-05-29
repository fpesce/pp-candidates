package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var wordlist string
var separator string
var maxWords int
var maxLength int

func init() {
	flag.StringVar(&wordlist, "w", "", "wordlist to generate pass from")
	flag.StringVar(&separator, "s", "", "separator between words")
	flag.IntVar(&maxWords, "n", 4, "maximum number of words in the pass phrase")
	flag.IntVar(&maxLength, "l", 64, "maximum length of the generated passphrase")
}

func recurse(prefix string, words []string, depth int) {
	if depth > 0 && len(prefix) <= maxLength {
		for _, w := range words {
			if depth == 1 {
				output := prefix + separator + w
				if len(output) <= maxLength {
					fmt.Println(output)
				}
			} else {
				var currentPrefix string
					if prefix != "" {
						currentPrefix = prefix + separator + w
					} else {
						currentPrefix = w
					}
				recurse(currentPrefix, words, depth - 1)
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
	recurse("", words, maxWords)
}
