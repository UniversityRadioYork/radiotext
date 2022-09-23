package data

import (
	"fmt"

	"github.com/UniversityRadioYork/myradio-go"
)

var (
	ErrNoNowPlaying = fmt.Errorf("nothing now playing")
)

func (e *Env) OutputNowPlaying() error {
	selectorInfo, err := e.MyRadioSession.GetSelectorInfo()
	if err != nil {
		return err
	}

	nowPlaying, err := e.MyRadioSession.GetNowPlaying(selectorInfo.Studio == myradio.SelectorOffAir)

	if nowPlaying.ID == 0 {
		return ErrNoNowPlaying
	}

	if err != nil {
		return err
	}

	e.SSHSession.OutputRadioTextMessage(fmt.Sprintf("Now Playing: %v - %v", nowPlaying.Title, nowPlaying.Artist))

	return nil
}
