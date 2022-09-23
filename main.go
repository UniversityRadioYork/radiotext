/** University Radio York - FM Radiotext

Author: Michael Grace <michael.grace@ury.org.uk>

**/

package main

import (
	"errors"
	"os"

	"github.com/UniversityRadioYork/myradio-go"
	"gopkg.in/yaml.v2"

	"github.com/UniversityRadioYork/radiotext/pkg/common"
	"github.com/UniversityRadioYork/radiotext/pkg/data"
	"github.com/UniversityRadioYork/radiotext/pkg/ssh"
)

func main() {
	// Load Config
	var config common.Config

	configFile, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(configFile, &config); err != nil {
		panic(err)
	}

	env := data.Env{}

	// Create MyRadio Session
	env.MyRadioSession, err = myradio.NewSessionFromKeyFile()
	if err != nil {
		panic(err)
	}

	// Create SSH Session
	env.SSHSession, err = ssh.OpenSSHConnection(config)
	if err != nil {
		panic(err)
	}
	defer env.SSHSession.Close()

	// Loop Over Entries and Output
	for {
		// Default Message
		env.SSHSession.OutputRadioTextMessage(config.DefaultMessage)

		// Custom Message
		// TODO

		// Now Playing
		if err := env.OutputNowPlaying(); err != nil {
			if !errors.Is(err, data.ErrNoNowPlaying) {
				env.SSHSession.OutputRadioTextMessage(config.DefaultMessage)
			}
		}

		// On Air Show
		if err := env.OutputOnAirShow(); err != nil {
			if !errors.Is(err, data.ErrNoShow) {
				env.SSHSession.OutputRadioTextMessage(config.DefaultMessage)
			}
		}

		// Future TODO:
		// On Air Next
		// URY News

	}
}
