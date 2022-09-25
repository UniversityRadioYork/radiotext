package data

import (
	"fmt"
	"log"
	"time"

	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/radiotext/pkg/common"
	"github.com/UniversityRadioYork/radiotext/pkg/ssh"
)

type RadiotextSession struct {
	SSHSession     *ssh.SSHSession
	MyRadioSession *myradio.Session

	OutputMessage     string
	PriorityWriteLock bool
}

func (s *RadiotextSession) OutputRadioTextMessage(msg string, highPriority bool) {
	if !highPriority {
		// if we're in a high priority write lock, and we get
		// a low priority message, we'll just store it so we
		// can refer to it later once the high priority message is over
		s.OutputMessage = msg

		if s.PriorityWriteLock {
			return
		}
	}

	for _, output := range common.SplitMessageToLength(msg, s.SSHSession.Config.MaxTextLength) {
		log.Println(output)
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

		time.Sleep(time.Duration(s.SSHSession.Config.WaitTime) * time.Second)
	}
}
