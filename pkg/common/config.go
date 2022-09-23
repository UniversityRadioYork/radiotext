package common

type Config struct {
	DefaultMessage string `yaml:"defaultMessage"`
	WaitTime       int    `yaml:"waitTime"`
	MaxTextLength  int    `yaml:"maxTextLength"`

	SSHHost string `yaml:"SSHHost"`
	SSHUser string `yaml:"SSHUser"`
	SSHPass string `yaml:"SSHPass"`
}
