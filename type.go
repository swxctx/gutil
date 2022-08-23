package gutil

import (
	"unicode"
	"unicode/utf8"
)

// IsExportedName is this an exported - upper case - name?
func IsExportedName(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}
