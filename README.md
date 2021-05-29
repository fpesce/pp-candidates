```
$ pp-candidates -h
Usage of ./pp-candidates:
  -l int
        maximum length of the generated passphrase (default 64)
  -n int
        maximum number of words in the pass phrase (default 4)
  -s string
        separator between words
  -w string
        wordlist to generate pass from


$ pp-candidates -n 4 -l 20 -w test.txt
```
