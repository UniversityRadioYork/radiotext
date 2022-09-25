package data

import (
	"fmt"
)

var (
	ErrNoShow = fmt.Errorf("not a show")
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

	s.OutputRadioTextMessage(fmt.Sprintf("On Air: %v with %v", currentShow.Title, currentShow.Presenters))

	return nil
}
