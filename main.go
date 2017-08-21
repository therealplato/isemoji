package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	rx      *regexp.Regexp
	err     error
	ok      bool
	results []string
	s       string
	n       int
)

func main() {
	rx, err = regexp.Compile(emojiRx)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n++
		s = scanner.Text()
		ok = rx.MatchString(s)
		if ok {
			results = append(results, s)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if len(results) == 0 {
		os.Stderr.WriteString("Found no emoji\n")
		os.Exit(1)
	}
	fmt.Printf("%v total lines of input, %v lines contained emoji:\n", n, len(results))
	for _, s := range results {
		fmt.Println(s)
	}
}
