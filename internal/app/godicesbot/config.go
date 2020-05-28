package godicesbot

// Config main bot configuration
type Config struct {
	AuthToken   string `toml:"auth_token"`
	TelegramURL string `toml:"telegram_url"`
}

// NewConfig creates new configuration
func NewConfig() *Config {
	return &Config{}
}
