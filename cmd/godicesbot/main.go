package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
	"github.com/sergonas/godicesbot/internal/app/godicesbot"
)

var (
	configPath string
	token      string
)

func init() {
	flag.StringVar(&configPath, "configPath", "config/dicesbot.toml", "path to toml config file")
	flag.StringVar(&token, "token", "NONE", "telegram bot-token")
}

func main() {
	flag.Parse()

	config := godicesbot.NewConfig(token)
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	bot := godicesbot.NewBotClient(*config)
	bot.ListenAndServe(func(update godicesbot.Update) {
		log.Info(update)
	})
}
