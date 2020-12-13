package main

import (
	"fmt"
	"os"
	"strings"
)

const MODE_ENCODE string = "encode"
const MODE_DECODE string = "decode"

var encodeTable = map[rune]rune{
	'a': '1',
	'e': '2',
	'i': '3',
	'o': '4',
	'u': '5',
}

var decodeTable = map[rune]rune{
	'1': 'a',
	'2': 'e',
	'3': 'i',
	'4': 'o',
	'5': 'u',
}

func decode(s string) string {
	var result strings.Builder

	for _, char := range s {
		if rChar, ok := decodeTable[char]; ok == true {
			result.WriteString(string(rChar))
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func encode(s string) string {
	var result strings.Builder

	for _, char := range s {
		if rChar, ok := encodeTable[char]; ok == true {
			result.WriteRune(rChar)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	var result string
	input := "hello world"
	mode := MODE_ENCODE

	if len(os.Args) >= 2 {
		input = os.Args[1]
		mode = os.Args[2]
	}

	switch mode {
	case MODE_ENCODE:
		result = encode(input)
	case MODE_DECODE:
		result = decode(input)
	}

	fmt.Printf("Original: %s\n", input)
	fmt.Printf("Result: %s\n", result)
}
