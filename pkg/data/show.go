package data

import (
	"fmt"
	"time"
)

var (
	ErrNoShow     = fmt.Errorf("not a show")
	ErrNoShowSoon = fmt.Errorf("no show soon")
)

func (s *RadiotextSession) OutputOnAirShow() error {
	currentAndNext, err := s.MyRadioSession.GetCurrentAndNext()
	if err != nil {
		return err
	}

	currentShow := currentAndNext.Current

	if currentShow.Id == 0 {
		return ErrNoShow
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

	if nextShow.Id == 0 || nextShow.StartTime.Before(
		time.Now().Add(time.Duration(10)*time.Minute),
	) {
		return ErrNoShowSoon
	}

	s.OutputRadioTextMessage(
		fmt.Sprintf("Coming Up: %v with %v", nextShow.Title, nextShow.Presenters),
		false,
	)

	return nil
}
