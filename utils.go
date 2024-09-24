package ygocards

import "unicode"

// UpperCamelをLowerSnakeに変換する.
func toLowerSnake(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_', unicode.ToLower(r))
		} else {
			result = append(result, unicode.ToLower(r))
		}
	}
	return string(result)
}
