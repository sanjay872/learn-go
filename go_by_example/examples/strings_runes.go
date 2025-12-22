package examples

import (
	"fmt"
	"unicode/utf8"
)

func StringsRunes() {
	const s = "Hello, 世界"

	fmt.Println("String:", s)

	// Length in bytes
	fmt.Println("Length in bytes:", len(s))

	// Iterate over bytes
	fmt.Println("Bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	// Iterate over runes
	fmt.Println("Runes:")
	for _, r := range s {
		fmt.Printf("%#U ", r)
	}

	fmt.Println()

	// Decode runes using utf8 package
	fmt.Println("Runes (using utf8 package):")
	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:]) // Decode rune starting at byte index i
		fmt.Printf("%#U ", r)
		w = width
	}

	fmt.Println()
	// Count runes
	runeCount := utf8.RuneCountInString(s)
	fmt.Println("Number of runes:", runeCount)
}
