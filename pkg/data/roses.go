package data

import (
	"fmt"
	"encoding/json"
	"net/http"
)

var ErrNoRosesLive = fmt.Errorf("no live roses")

type stream struct {
	id string `json:"id"`
	live bool `json:"live"`
	event string `json:"event"`
}

func (s *RadiotextSession) OutputRosesData(rosesAPI string) error {
	res, err := http.Get(rosesAPI)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var data []stream
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	var roseslive bool
	for _, strm := range data {
		if strm.id == "Broadcast" && strm.live {
			s.OutputRadioTextMessage(strm.event, false)
			roseslive = true
			continue
		}
	}

	if !roseslive {
		return ErrNoRosesLive
	}

}