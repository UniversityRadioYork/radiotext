package data

import (
	"fmt"
	"strings"
	"time"
)

func withInTitle(title string) bool {
	for _, word := range strings.Split(title, " ") {
		if word == "with" {
			return true
		}
	}

	return false
}

func (s *RadiotextSession) OutputOnAirShow() error {
	currentAndNext, err := s.MyRadioSession.GetCurrentAndNext()
	if err != nil {
		return err
	}

	currentShow := currentAndNext.Current

	if currentShow.Id == 0 {
		// No Current Showw
		return nil
	}

	onAirMessage := fmt.Sprintf("On Air: %v", currentShow.Title)
	if !withInTitle(currentShow.Title) {
		onAirMessage = fmt.Sprintf("%v with %v", onAirMessage, currentShow.Presenters)
	}

	s.OutputRadioTextMessage(
		onAirMessage,
		false,
	)

	return nil
}

func (s *RadiotextSession) NextShowHandler() error {
	currentAndNext, err := s.MyRadioSession.GetCurrentAndNext()
	if err != nil {
		return err
	}

	nextShow := currentAndNext.Next

	if nextShow.Id == 0 || nextShow.StartTime.After(
		time.Now().Add(time.Duration(10)*time.Minute),
	) {
		// Not going to show it yet
		return nil
	}

	comingUpMessage := fmt.Sprintf("Coming Up: %v", nextShow.Title)
	if !withInTitle(nextShow.Title) {
		comingUpMessage = fmt.Sprintf("%v with %v", comingUpMessage, nextShow.Presenters)
	}

	s.OutputRadioTextMessage(
		comingUpMessage,
		false,
	)

	return nil
}
