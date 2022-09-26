package data

import (
	"fmt"
	"time"
)

var hourNumsToWords = map[int]string{
	0:  "Twelve", // 12 % 12 = 0
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
}

func (s *RadiotextSession) URYNewsHandler() {
	// we'll wait until we're less than 2 mins past
	// then send a high priority message of URY news

	var isNews bool
	for {
		now := time.Now()

		if !isNews {
			if now.Minute() < 2 {
				isNews = true
				s.PriorityWriteLock = true
				s.OutputRadioTextMessage(
					fmt.Sprintf("URY News at %v", hourNumsToWords[now.Hour()%12]), true)
			}
		} else {
			if now.Minute() >= 2 {
				isNews = false
				s.PriorityWriteLock = false
				s.OutputRadioTextMessage(
					s.OutputMessage,
					false,
				)
			}
		}

		time.Sleep(time.Duration(15) * time.Second)
	}
}
