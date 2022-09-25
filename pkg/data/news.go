package data

import (
	"fmt"
	"time"
)

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
					fmt.Sprintf("URY News at %v", now.Format("3")), true)
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
