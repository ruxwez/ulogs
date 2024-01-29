package ulogs

import (
	"errors"
	"time"

	"github.com/gtuk/discordwebhook"
)

var (
	errToConnectWithDiscord = errors.New("error to connect with discord")
)

type discordRepository struct {
	config *Config
}

func (s *discordRepository) sendDiscordLog(typeLog string, content string) error {
	var username = "ULogs"
	var url = s.config.DiscordURL
	var color = "0x000000"
	var TimeNow = time.Now().Format("2006-01-02 15:04:05")

	if typeLog == TYPE_INFO {
		color = "5814783"
	}

	if typeLog == TYPE_ERROR {
		color = "16738665"
	}

	if typeLog == TYPE_WARNING {
		color = "16767849"
	}

	embed := discordwebhook.Embed{
		Title:       &TimeNow,
		Description: &content,
		Color:       &color,
		Author: &discordwebhook.Author{
			Name: &typeLog,
		},
		Footer: &discordwebhook.Footer{
			Text: &s.config.Name,
		},
	}

	message := discordwebhook.Message{
		Username: &username,
		Embeds:   &[]discordwebhook.Embed{embed},
	}

	err := discordwebhook.SendMessage(url, message)

	if err != nil {
		return errToConnectWithDiscord
	}

	return nil
}
