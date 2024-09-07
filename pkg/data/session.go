package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/radiotext/pkg/common"
	"github.com/UniversityRadioYork/radiotext/pkg/ssh"
)

type outputState int

const (
	OutputRun outputState = iota
	OutputPause
	OutputStop
)

var mostRecentMessage string

func (s *RadiotextSession) checkOutputState() outputState {
	res, err := http.Get(s.SSHSession.Config.EndpointForRunningCheck)
	if err != nil {
		log.Println(err.Error())
		return OutputRun
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return OutputRun
	}

	var state map[string]int
	if err = json.Unmarshal(body, &state); err != nil {
		log.Println(err.Error())
		return OutputRun
	}

	return []outputState{OutputRun, OutputPause, OutputStop}[state[s.SSHSession.Config.RunningCheckKey]]

}

type RadiotextSession struct {
	SSHSession     *ssh.SSHSession
	MyRadioSession *myradio.Session

	OutputMessage     string
	PriorityWriteLock bool
}

func (s *RadiotextSession) OutputRadioTextMessage(msg string, highPriority bool) {
	var fullyOutputed bool

	defer func() {
		if !fullyOutputed {
			time.Sleep(time.Duration(s.SSHSession.Config.WaitTime) * time.Second)
		}
	}()

	if !highPriority {
		// if we're in a high priority write lock, and we get
		// a low priority message, we'll just store it so we
		// can refer to it later once the high priority message is over
		s.OutputMessage = msg

		if s.PriorityWriteLock {
			return
		}
	}

	m, err := common.ToAscii(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg = m

	for _, output := range common.SplitMessageToLength(msg, s.SSHSession.Config.MaxTextLength) {
		switch s.checkOutputState() {
		case OutputRun:
			// run
		case OutputPause:
			time.Sleep(time.Duration(s.SSHSession.Config.WaitTime/3) * time.Second)
			return
		case OutputStop:
			if mostRecentMessage == s.SSHSession.Config.DefaultMessage {
				time.Sleep(time.Duration(s.SSHSession.Config.WaitTime/3) * time.Second)
				return
			}
			if output != s.SSHSession.Config.DefaultMessage {
				s.OutputRadioTextMessage(s.SSHSession.Config.DefaultMessage, false)
				return
			}
		}

		log.Println(output)

		if s.SSHSession.DoConnection {
			_, err := s.SSHSession.Stdin.Write([]byte(
				fmt.Sprintf(
					"echo -ne \"%v\\f\" > %v\n",
					output,
					s.SSHSession.Config.OutputDevice,
				),
			))

			if err != nil {
				if err.Error() == "EOF" {
					s.SSHSession.Close()
					s.SSHSession.CreateConnection()
					s.OutputRadioTextMessage(msg, highPriority)
					return
				}

				fmt.Println(err)
			}

		}

		mostRecentMessage = output

		time.Sleep(time.Duration(s.SSHSession.Config.WaitTime) * time.Second)
	}

	fullyOutputed = true
}
