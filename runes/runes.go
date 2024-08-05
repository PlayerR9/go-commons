package runes

import "unicode"

// ToInt converts a rune to an integer if possible. Conversion is case-insensitive and
// values from 0-9 and a-z are converted to 0-35.
//
// Parameters:
//   - char: The rune to convert.
//
// Returns:
//   - int: The converted integer.
//   - bool: True if the conversion was successful. False otherwise.
//
// Example:
//
//	digit, ok := ToInt('A')
//	if !ok {
//		panic("Could not convert 'A' to an integer")
//	}
//
//	fmt.Println(digit) // 10
func ToInt(char rune) (int, bool) {
	ok := unicode.IsDigit(char)
	if ok {
		return int(char - '0'), true
	}

	ok = unicode.IsLetter(char)
	if !ok {
		return 0, false
	}

	char = unicode.ToLower(char)

	return int(char - 'a' + 10), true
}
