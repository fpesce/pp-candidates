```
$ pp-candidates -h
Usage of ./pp-candidates:
  -a    avoid phrase with twice the same word
  -e string
        add a last string/char to every phrases
  -l int
        maximum length of the generated passphrase (default 64)
  -m int
        minimum number of words in the pass phrase (default 4)
  -n int
        maximum number of words in the pass phrase (default 4)
  -s string
        separator between words
  -w string
        wordlist to generate pass from

$ ./pp-candidates -w test.txt -a -s " " -l 43 -n 8 -e "!"
```
