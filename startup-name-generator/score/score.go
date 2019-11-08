package score

import (
	"github.com/NandanSatheesh/go-lang-playground/startup-name-generator/constants"
	"github.com/NandanSatheesh/go-lang-playground/startup-name-generator/utils"
	"math/rand"
	str "strings"
	"time"
)

func Score(word string) float32 {
	rand.Seed(time.Now().Unix())
	var score float32
	score = 0
	score += ScoreLength(word)
	score += ScoreSuffix(word)
	score += ScoreSyllables(word)
	score += 0.4 + rand.Float32()
	return score
}

func ScoreSuffix(word string) float32 {
	actualSuffixLength := len(constants.ACTUAL_SUFFIX)

	if str.Contains(word, constants.ACTUAL_SUFFIX[rand.Intn(actualSuffixLength)]) {
		return -1.5
	} else {
		return 0
	}
}

func ScoreSyllables(word string) float32 {
	var syllables float32
	syllables = float32(utils.CountSyllables([]byte(word)))

	switch {
	case syllables == 2:
		return 6.1
	case syllables == 3:
		return 6
	case syllables > 4:
		return 2
	default:
		return 4
	}
}

func ScoreLength(word string) float32 {
	normalizedLength := len(utils.Normalize(word))
	if normalizedLength < 9 {
		return 0.1
	} else {
		return 0
	}
}
