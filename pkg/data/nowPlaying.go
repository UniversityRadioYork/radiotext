package data

import (
	"fmt"

	"github.com/UniversityRadioYork/myradio-go"
)

var (
	ErrNoNowPlaying = fmt.Errorf("nothing now playing")
)

func (s *RadiotextSession) OutputNowPlaying() error {
	selectorInfo, err := s.MyRadioSession.GetSelectorInfo()
	if err != nil {
		return err
	}

	nowPlaying, err := s.MyRadioSession.GetNowPlaying(selectorInfo.Studio == myradio.SelectorOffAir)

	if nowPlaying.ID == 0 {
		return ErrNoNowPlaying
	}

	if err != nil {
		return err
	}

	s.OutputRadioTextMessage(fmt.Sprintf("Now Playing: %v - %v", nowPlaying.Title, nowPlaying.Artist))

	return nil
}
