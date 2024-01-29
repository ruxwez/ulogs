package ulogs

type Config struct {
	// Name of the service
	Name string

	// Prefix to put on the Logs File
	FilePrefix string

	// DISCORD LOGS CONFIG
	// Discord Webhook URL
	DiscordURL string
}
