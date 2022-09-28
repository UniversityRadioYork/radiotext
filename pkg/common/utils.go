package common

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func SplitMessageToLength(msg string, length int) []string {
	var splits []string

	var subMsg string
	for _, word := range strings.Split(msg, " ") {
		var checkMsg string
		if len(subMsg) == 0 {
			checkMsg = word
		} else {
			checkMsg = fmt.Sprintf("%s %s", subMsg, word)
		}

		if len(checkMsg) <= length {
			subMsg = checkMsg
		} else {
			splits = append(splits, subMsg)
			subMsg = word
		}
	}
	splits = append(splits, subMsg)
	return splits
}

func ToAscii(str string) (string, error) {
	result, _, err := transform.String(
		transform.Chain(
			norm.NFD,
			runes.Remove(runes.In(unicode.Mn))),
		str)

	if err != nil {
		return "", err
	}

	return result, nil
}
