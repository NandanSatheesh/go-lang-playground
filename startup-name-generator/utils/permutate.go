package utils

import "github.com/NandanSatheesh/go-lang-playground/startup-name-generator/constants"

func Permutate(words []string) []string {
	if len(words) == 0 {
		return PermuteFixes(constants.PREFIXES, constants.SUFFIXES)
	} else {
		return append(PermuteFixes(constants.PREFIXES, words), PermuteFixes(words, constants.SUFFIXES)...)
	}
}

func PermuteFixes(prefixes []string, suffixes []string) []string {
	var listOfNames []string

	for _, prefix := range prefixes {
		for _, suffix := range suffixes {
			if prefix != suffix {
				listOfNames = append(listOfNames, prefix+" "+suffix)
			}
		}
	}

	return listOfNames
}
