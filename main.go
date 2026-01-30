package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var enchantmentTableFont = map[rune]string{
	'a': "ᔑ", 'b': "ʖ", 'c': "ᓵ", 'd': "↸", 'e': "ᒷ", 'f': "⎓", 'g': "⊣", 'h': "⍑",
	'i': "╎", 'j': "⋮", 'k': "ꖌ", 'l': "ꖎ", 'm': "ᒲ", 'n': "リ", 'o': "𝙹", 'p': "!",
	'q': "¡", 'r': "ᑑ", 's': "∷", 't': "ᓭ", 'u': "ℸ", 'v': "⚍", 'w': "⍊", 'x': "∴",
	'y': "̇/", 'z': "||",
	'A': "ᔑ", 'B': "ʖ", 'C': "ᓵ", 'D': "↸", 'E': "ᒷ", 'F': "⎓", 'G': "⊣", 'H': "⍑",
	'I': "╎", 'J': "⋮", 'K': "ꖌ", 'L': "ꖎ", 'M': "ᒲ", 'N': "リ", 'O': "𝙹", 'P': "!",
	'Q': "¡", 'R': "ᑑ", 'S': "∷", 'T': "ᓭ", 'U': "ℸ", 'V': "⚍", 'W': "⍊", 'X': "∴",
	'Y': "̇/", 'Z': "||",
	'0': "0", '1': "1", '2': "2", '3': "3", '4': "4", '5': "5", '6': "6", '7': "7", '8': "8", '9': "9",
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
		line := scanner.Text()
		fmt.Println(translateToEnchantmentFont(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
