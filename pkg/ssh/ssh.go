package ssh

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/UniversityRadioYork/radiotext/pkg/common"
)

type SSHSession struct {
	config  common.Config
	conn    *ssh.Client
	session *ssh.Session
	stdin   io.WriteCloser
}

func OpenSSHConnection(config common.Config) (*SSHSession, error) {
	conf := &ssh.ClientConfig{
		User:            config.SSHUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password(config.SSHPass)},
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", config.SSHHost), conf)
	if err != nil {
		return nil, err
	}

	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		return nil, err
	}

	if err = session.Shell(); err != nil {
		return nil, err
	}

	return &SSHSession{
		config:  config,
		conn:    conn,
		session: session,
		stdin:   stdin,
	}, nil
}

func (s *SSHSession) Close() {
	s.conn.Close()
	s.session.Close()
}

func splitMessageToLength(msg string, length int) []string {
	var splits []string

	var subMsg string
	for _, word := range strings.Split(msg, " ") {
		var checkMsg string
		if len(subMsg) == 0 {
			checkMsg = word
		} else {
			checkMsg = fmt.Sprintf("%s %s", subMsg, word)
		}

		if len(checkMsg) <= length {
			subMsg = checkMsg
		} else {
			splits = append(splits, subMsg)
			subMsg = word
		}
	}
	splits = append(splits, subMsg)
	return splits
}

func (s *SSHSession) OutputRadioTextMessage(msg string) {
	for _, output := range splitMessageToLength(msg, s.config.MaxTextLength) {
		log.Println(output)
		_, err := s.stdin.Write([]byte(fmt.Sprintf("echo -ne \"%v\\f\" > /dev/ttyUSB0\n", output)))
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Duration(s.config.WaitTime) * time.Second)
	}
}
