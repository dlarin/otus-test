package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (res string, _ error) {
	if len(s) == 0 {
		return "", nil
	}

	runeArr := []rune(s)

	if unicode.IsDigit(runeArr[0]) {
		return "", ErrInvalidString
	}

	var result strings.Builder

	isEscaping := false
	for i := 0; i <= len(runeArr)-1; i++ {
		if i == len(runeArr)-1 {
			if string(runeArr[i]) == `\` {
				return "", ErrInvalidString
			}

			result.WriteRune(runeArr[i])
			continue
		}

		nextRune := runeArr[i+1]

		// Escaping
		if string(runeArr[i]) == `\` && !isEscaping {
			if unicode.IsDigit(nextRune) || string(nextRune) == `\` {
				isEscaping = true
				continue
			}

			return "", ErrInvalidString
		}

		if (isCopyRune(runeArr, i) || isEscaping) && unicode.IsDigit(nextRune) {
			count, _ := strconv.Atoi(string(nextRune))
			result.WriteString(
				strings.Repeat(string(runeArr[i]), count),
			)
			i++
			continue
		}

		if unicode.IsDigit(runeArr[i]) && unicode.IsDigit(runeArr[i-1]) {
			return "", ErrInvalidString
		}
		result.WriteRune(runeArr[i])

		isEscaping = false
	}

	return result.String(), nil
}

func isCopyRune(r []rune, i int) bool {
	return unicode.IsLetter(r[i]) || unicode.IsSpace(r[i])
}
