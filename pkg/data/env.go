package data

import (
	"github.com/UniversityRadioYork/myradio-go"
	"github.com/UniversityRadioYork/radiotext/pkg/ssh"
)

type Env struct {
	SSHSession     *ssh.SSHSession
	MyRadioSession *myradio.Session
}
