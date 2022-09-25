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

	var session data.RadiotextSession

	// Create MyRadio Session
	session.MyRadioSession, err = myradio.NewSession(config.APIKey)
	if err != nil {
		panic(err)
	}

	// Create SSH Session
	session.SSHSession, err = ssh.OpenSSHConnection(config)
	if err != nil {
		panic(err)
	}
	defer session.SSHSession.Close()

	// URY News
	go session.URYNewsHandler()

	// Loop Over Entries and Output
	for {
		// Default Message
		session.OutputRadioTextMessage(config.DefaultMessage, false)

		// Custom Message
		// TODO

		// Now Playing
		if err := session.OutputNowPlaying(); err != nil {
			if !errors.Is(err, data.ErrNoNowPlaying) {
				session.OutputRadioTextMessage(config.DefaultMessage, false)
			}
		}

		// On Air Show
		if err := session.OutputOnAirShow(); err != nil {
			if !errors.Is(err, data.ErrNoShow) {
				session.OutputRadioTextMessage(config.DefaultMessage, false)
			}
		}

		// Future TODO:
		// On Air Next

	}
}
