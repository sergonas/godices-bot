package godicesbot

// Config main bot configuration
type Config struct {
	AuthToken   string
	TelegramURL string `toml:"telegram_url"`
}

// NewConfig creates new configuration
func NewConfig(token string) *Config {
	return &Config{AuthToken: token}
}
