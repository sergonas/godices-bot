package godicesbot

import "fmt"

// SendAnnouncment send announcment to all event users
func SendAnnouncment(text string) error {
	fmt.Println("Announcment: " + text)
	return nil
}
