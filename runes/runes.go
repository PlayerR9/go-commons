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

// FromInt converts an integer to a rune if possible. Conversion is case-insensitive and
// values from 0-9 and a-z are converted to 0-35.
//
// Parameters:
//   - digit: The integer to convert.
//
// Returns:
//   - rune: The converted rune.
//   - bool: True if the conversion was successful. False otherwise.
//
// Example:
//
//	char, ok := FromInt(10)
//	if !ok {
//		panic("Could not convert 10 to a rune")
//	}
//
//	fmt.Println(char) // 'A'
func FromInt(digit int) (rune, bool) {
	if digit < 0 || digit > 35 {
		return 0, false
	}

	if digit < 10 {
		return rune(digit + '0'), true
	}

	return rune(digit - 10 + 'a'), true
}
