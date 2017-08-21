package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	var tcs = []string{
		"hello",
		"🌏",
		"your weather is 🌤 today",
		"615 °C",
		"born in the 🇺🇸",
		"asdf",
	}
	var expected = []bool{
		false, true, true, true, true, false,
	}

	rx, err = regexp.Compile(emojiRx)
	for i, tc := range tcs {
		actual := rx.Match([]byte(tc))
		assert.Equal(t, expected[i], actual)
		fmt.Printf("%s:\n%v\n\n", tc, found)
	}

}
