package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var enchantmentTableFont = map[rune]string{
	'a': "ᔑ", 'b': "ʖ", 'c': "ᓵ", 'd': "↸", 'e': "ᒷ", 'f': "⎓",
	'g': "⊣", 'h': "⍑", 'i': "╎", 'j': "⋮", 'k': "ꖌ", 'l': "ꖎ",
	'm': "ᒲ", 'n': "リ", 'o': "𝙹", 'p': "ᓭ", 'q': "¡", 'r': "ᑑ",
	's': "∷", 't': "ℸ", 'u': "⚍", 'v': "⍊", 'w': "∴", 'x': "⨅",
	'y': "॥", 'z': "ᔑ",

	'A': "ᔑ", 'B': "ʖ", 'C': "ᓵ", 'D': "↸", 'E': "ᒷ", 'F': "⎓",
	'G': "⊣", 'H': "⍑", 'I': "╎", 'J': "⋮", 'K': "ꖌ", 'L': "ꖎ",
	'M': "ᒲ", 'N': "リ", 'O': "𝙹", 'P': "ᓭ", 'Q': "¡", 'R': "ᑑ",
	'S': "∷", 'T': "ℸ", 'U': "⚍", 'V': "⍊", 'W': "∴", 'X': "⨅",
	'Y': "॥", 'Z': "ᔑ",

	'0': "0", '1': "1", '2': "2", '3': "3", '4': "4",
	'5': "5", '6': "6", '7': "7", '8': "8", '9': "9",
}

func translateToEnchantmentFont(text string) string {
	var result []rune
	for _, char := range text {
		if translated, exists := enchantmentTableFont[char]; exists {
			result = append(result, []rune(translated)...)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println(translateToEnchantmentFont(strings.Join(os.Args[1:], " ")))
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(translateToEnchantmentFont(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
