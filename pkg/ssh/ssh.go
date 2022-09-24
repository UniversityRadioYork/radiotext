package ssh

import (
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/UniversityRadioYork/radiotext/pkg/common"
)

type SSHSession struct {
	config  common.Config
	sshConf ssh.ClientConfig
	conn    *ssh.Client
	session *ssh.Session
	stdin   io.WriteCloser
}

func OpenSSHConnection(config common.Config) (*SSHSession, error) {
	s := SSHSession{
		config: config,
		sshConf: ssh.ClientConfig{
			User:            config.SSHUser,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth:            []ssh.AuthMethod{ssh.Password(config.SSHPass)},
		},
	}

	if err := s.createConnection(); err != nil {
		return nil, err
	}

	return &s, nil

}

func (s *SSHSession) Close() {
	s.conn.Close()
	s.session.Close()
}

func (s *SSHSession) createConnection() error {
	fmt.Println("connecting SSH")

	var err error
	s.conn, err = ssh.Dial("tcp", fmt.Sprintf("%s:22", s.config.SSHHost), &s.sshConf)
	if err != nil {
		return err
	}

	s.session, err = s.conn.NewSession()
	if err != nil {
		return err
	}

	s.stdin, err = s.session.StdinPipe()
	if err != nil {
		return err
	}

	if err = s.session.Shell(); err != nil {
		return err
	}

	return nil

}

func (s *SSHSession) OutputRadioTextMessage(msg string) {
	for _, output := range common.SplitMessageToLength(msg, s.config.MaxTextLength) {
		log.Println(output)
		_, err := s.stdin.Write([]byte(fmt.Sprintf("echo -ne \"%v\\f\" > /dev/ttyUSB0\n", output)))
		if err != nil {
			if err.Error() == "EOF" {
				s.createConnection()
				s.OutputRadioTextMessage(msg)
				return
			}

			fmt.Println(err)
		}

		time.Sleep(time.Duration(s.config.WaitTime) * time.Second)
	}
}
