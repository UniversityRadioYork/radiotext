package data

import (
	"fmt"

	"github.com/UniversityRadioYork/myradio-go"
)

func (s *RadiotextSession) OutputNowPlaying() error {
	selectorInfo, err := s.MyRadioSession.GetSelectorInfo()
	if err != nil {
		return err
	}

	nowPlaying, err := s.MyRadioSession.GetNowPlaying(selectorInfo.Studio == myradio.SelectorOffAir)

	if nowPlaying.ID == 0 {
		// No Now Playing Info
		return nil
	}

	if err != nil {
		return err
	}

	s.OutputRadioTextMessage(
		fmt.Sprintf("Now Playing: %v - %v", nowPlaying.Title, nowPlaying.Artist),
		false,
	)

	return nil
}
