package ulogs

import "fmt"

type Ulogs interface {
	Send(typeLog string, message string)
}

type ulogs struct {
	config *Config
}

func NewInstance(c *Config) Ulogs {
	return &ulogs{
		config: c,
	}
}

func (u *ulogs) Send(typeLog string, message string) {

	_fileInstance := filesRepository{
		config: u.config,
	}
	// Add the message on the logs file.
	err := _fileInstance.addLogFile(typeLog, message)

	if err != nil {
		fmt.Println(err.Error())
	}

	if u.config.DiscordURL == "" {
		return
	}

	_discordInstance := discordRepository{
		config: u.config,
	}

	err = _discordInstance.sendDiscordLog(typeLog, message)

	if err != nil {
		fmt.Println(err.Error())
	}
}
