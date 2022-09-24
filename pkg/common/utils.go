package common

import (
	"fmt"
	"strings"
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
