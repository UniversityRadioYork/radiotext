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
}

func (s *RadiotextSession) OutputRadioTextMessage(msg string) {
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
				s.OutputRadioTextMessage(msg)
				return
			}

			fmt.Println(err)
		}

		time.Sleep(time.Duration(s.SSHSession.Config.WaitTime) * time.Second)
	}
}
