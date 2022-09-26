package data

import (
	"fmt"
	"time"
)

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

	s.OutputRadioTextMessage(
		fmt.Sprintf("On Air: %v with %v", currentShow.Title, currentShow.Presenters),
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

	s.OutputRadioTextMessage(
		fmt.Sprintf("Coming Up: %v with %v", nextShow.Title, nextShow.Presenters),
		false,
	)

	return nil
}
