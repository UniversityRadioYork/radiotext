package data

import (
	"fmt"
)

var (
	ErrNoShow = fmt.Errorf("not a show")
)

func (e *Env) OutputOnAirShow() error {
	currentAndNext, err := e.MyRadioSession.GetCurrentAndNext()
	if err != nil {
		return err
	}

	currentShow := currentAndNext.Current

	if currentShow.Id == 0 {
		return ErrNoShow
	}

	e.SSHSession.OutputRadioTextMessage(fmt.Sprintf("On Air: %v with %v", currentShow.Title, currentShow.Presenters))

	return nil
}
