package utilities

import (
	"fmt"
	"unicode"

	"golang.org/x/text/unicode/norm"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)

	if err != nil {
		fmt.Println("error transforming string")
	}

	return output
}
