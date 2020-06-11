package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/sergonas/godicesbot/internal/app/godicesbot"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "dicesbot.toml", "path to toml config file")
}

func main() {
	flag.Parse()

	config := godicesbot.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	bot := godicesbot.NewBotClient(*config)
	bot.ListenAndServe(func(update godicesbot.Update) {
		fmt.Printf("%v\n", update)
	})
}
