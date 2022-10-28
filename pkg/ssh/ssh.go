package ssh

import (
	"fmt"
	"io"

	"golang.org/x/crypto/ssh"

	"github.com/UniversityRadioYork/radiotext/pkg/common"
)

type SSHSession struct {
	Config  common.Config
	sshConf ssh.ClientConfig
	conn    *ssh.Client
	session *ssh.Session
	Stdin   io.WriteCloser
}

func OpenSSHConnection(config common.Config) (*SSHSession, error) {
	s := SSHSession{
		Config: config,
		sshConf: ssh.ClientConfig{
			User:            config.SSHUser,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth:            []ssh.AuthMethod{ssh.Password(config.SSHPass)},
		},
	}

	if err := s.CreateConnection(); err != nil {
		return nil, err
	}

	return &s, nil

}

func (s *SSHSession) Close() {
	s.conn.Close()
	s.session.Close()
}

func (s *SSHSession) CreateConnection() error {
	fmt.Println("connecting SSH")

	var err error
	s.conn, err = ssh.Dial("tcp", fmt.Sprintf("%s:22", s.Config.SSHHost), &s.sshConf)
	if err != nil {
		return err
	}

	s.session, err = s.conn.NewSession()
	if err != nil {
		return err
	}

	s.Stdin, err = s.session.StdinPipe()
	if err != nil {
		return err
	}

	if err = s.session.Shell(); err != nil {
		return err
	}

	if _, err := s.Stdin.Write([]byte(
		fmt.Sprintf(
			"sudo stty -F %s 9600 -ixoff -ixon -cread clocal",
			s.Config.OutputDevice,
		),
	)); err != nil {
		return err
	}

	return nil

}
