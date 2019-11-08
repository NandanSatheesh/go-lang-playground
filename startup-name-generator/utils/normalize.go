package utils

import str "strings"

func Normalize(word string) string {

	word = str.ReplaceAll(word, "e i", "i")
	word = str.ReplaceAll(word, "th t", "t")
	word = str.ReplaceAll(word, " ", "")
	word = str.ToUpper(word)
	return word
}
