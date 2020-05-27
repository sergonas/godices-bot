package main

import (
	"fmt"

	"github.com/sergonas/godicesbot/internal/app/godicesbot"
)

func init() {
	fmt.Println("Init method")
}

func main() {
	godicesbot.SendAnnouncment("WARNING")
}
