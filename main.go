package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	rx        *regexp.Regexp
	err       error
	found, ok bool
	results   []string
	s         string
)

func main() {
	rx, err = regexp.Compile(emojiRx)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s = scanner.Text()
		ok = rx.MatchString(s)
		if ok {
			found = true
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
	fmt.Println("Found emoji:")
	for _, s := range results {
		fmt.Println(s)
	}
}
