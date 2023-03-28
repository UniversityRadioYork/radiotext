package common

type Config struct {
	DefaultMessage string `yaml:"defaultMessage"`
	WaitTime       int    `yaml:"waitTime"`
	MaxTextLength  int    `yaml:"maxTextLength"`

	APIKey string `yaml:"apiKey"`

	SSHHost      string `yaml:"SSHHost"`
	SSHUser      string `yaml:"SSHUser"`
	SSHPass      string `yaml:"SSHPass"`
	OutputDevice string `yaml:"outputDevice"`

	RosesAPI string `yaml:"rosesAPI"`
}
