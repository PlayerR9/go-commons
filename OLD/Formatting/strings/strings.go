package strings

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/dustin/go-humanize"
)

// DateStringer prints the date in the format "1st January, 2006".
//
// Parameters:
//
//   - date: The date to print.
//
// Returns:
//
//   - string: The date in the format "1st January, 2006".
func DateStringer(date time.Time) string {
	return fmt.Sprintf("%s %v, %d",
		humanize.Ordinal(date.Day()),
		date.Month(),
		date.Year(),
	)
}

// TimeStringer prints the time in the format "3:04 PM".
//
// Parameters:
//
//   - time: The time to print.
//
// Returns:
//
//   - string: The time in the format "3:04 PM".
func TimeStringer(time time.Time) string {
	return time.Format("3:04 PM")
}

// StringsJoiner joins a list of fmt.Stringer values using a separator.
//
// Parameters:
//   - values: The list of fmt.Stringer values to join.
//   - sep: The separator to use when joining the strings.
//
// Returns:
//   - string: The string representation of the values.
func StringsJoiner[T fmt.Stringer](values []T, sep string) string {
	stringValues := make([]string, 0, len(values))

	for _, value := range values {
		stringValues = append(stringValues, value.String())
	}

	return strings.Join(stringValues, sep)
}

// ArrayFormatter formats a list of strings as an array.
//
// Parameters:
//   - values: The list of strings to format.
//
// Returns:
//   - string: The formatted array.
func ArrayFormatter(values []string) string {
	return "[" + strings.Join(values, ", ") + "]"
}

// FixTabStop fixes the tab stops in a string.
//
// The initial level and the integer return are used for
// chaining multiple calls to this function.
//
// Parameters:
//   - init: The initial level of the tab stop.
//   - tabSize: The size of the tab.
//   - spacing: The spacing to use for the tab stop.
//   - str: The string to fix the tab stops for.
//
// Returns:
//   - string: The string with the tab stops fixed.
//   - int: The total number of characters in the string.
//
// Behaviors:
//   - If the tabSize is less than 1, it is set to 1.
//   - If the initial tab stop is less than 0, it is set to 0.
//   - If the string is empty, an empty string is returned.
func FixTabStop(init, tabSize int, spacing, str string) (string, int) {
	var total int

	if init < 0 {
		total = 0
		init = 0
	} else {
		total = init
		init = tabSize - (init % tabSize)
	}

	if len(str) == 0 {
		return "", total
	}

	if tabSize < 1 {
		tabSize = 1
	}

	var builder strings.Builder

	if init > 0 {
		init = tabSize - (init % tabSize)

		for i := init; i > 0; i-- {
			r, size := utf8.DecodeRuneInString(str)
			str = str[size:]

			if r == '\t' {
				repStr := strings.Repeat(spacing, i)
				total += i

				builder.WriteString(repStr)
				break
			} else {
				total++
				builder.WriteRune(r)
			}

			if len(str) == 0 {
				return builder.String(), total
			}
		}
	}

	for {
		for i := tabSize; i > 0; i-- {
			r, size := utf8.DecodeRuneInString(str)
			str = str[size:]

			if r == '\t' {
				repStr := strings.Repeat(spacing, i)
				total += i

				builder.WriteString(repStr)
				break
			} else {
				total++
				builder.WriteRune(r)
			}

			if len(str) == 0 {
				return builder.String(), total
			}
		}
	}
}
