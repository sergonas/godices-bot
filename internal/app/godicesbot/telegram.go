package godicesbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// BotClient client to specified bot
type BotClient struct {
	config          Config
	lastMaxUpdateID int
}

// NewBotClient creates and initializes telegram bot client
func NewBotClient(configuration Config) *BotClient {
	return &BotClient{configuration, -20}
}

// ListenAndServe start bot handler
func (b *BotClient) ListenAndServe(onUpdate func(Update)) {
	for {
		updates := b.GetMessages()
		for _, update := range updates {
			onUpdate(update)
			b.updateLastMaxUpdate(update.UpdateID)
		}
	}
}

func (b *BotClient) updateLastMaxUpdate(maxUpdateID int) {
	if maxUpdateID >= b.lastMaxUpdateID {
		b.lastMaxUpdateID = maxUpdateID + 1
	}
}

// Response ...
type Response struct {
	IsOk   bool     `json:"ok"`
	Result []Update `json:"result"`
}

// Update This object represents an incoming update.
type Update struct {
	UpdateID          int      `json:"update_id"`
	Message           *Message `json:"message"`
	EditedMessage     *Message `json:"edited_message"`
	ChannelPost       *Message `json:"channel_post"`
	EditedChannelPost *Message `json:"edited_channel_post"`
}

// Message This object represents a message.
type Message struct {
	MessageID            int      `json:"message_id"`
	From                 *User    `json:"from"`
	Date                 int      `json:"date"`
	Chat                 Chat     `json:"chat"`
	ForwardFrom          *User    `json:"forward_from"`
	ForwardFromChat      *Chat    `json:"forward_from_chat"`
	ForwardFromMessageID int      `json:"forward_from_message_id"`
	ForwardSignature     string   `json:"forward_signature"`
	ForwardSenderName    string   `json:"forward_sender_name"`
	ForwardDate          int      `json:"forward_date"`
	ReplyToMessage       *Message `json:"reply_to_message"`
	ViaBot               *User    `json:"via_bot"`
	EditDate             int      `json:"edit_date"`
	MediaGroupID         int      `json:"media_group_id"`
	AuthorSignature      string   `json:"author_signature"`
	Text                 string   `json:"text"`
}

// User This object represents a Telegram user or bot.
type User struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// Chat This object represents a chat.
type Chat struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Username string `json:"username"`
}

// SendAnnouncment send announcment to all event users
func SendAnnouncment(text string) error {
	fmt.Println("Announcment: " + text)
	return nil
}

// GetMessages receives new messages from telegram server
func (b *BotClient) GetMessages() []Update {
	res, err := b.makeRequest("getUpdates", "{\"offset\": "+strconv.Itoa(b.lastMaxUpdateID))
	fmt.Printf("LastUpdateID = %v\n", b.lastMaxUpdateID)

	defer res.Body.Close()
	if err != nil {
		log.Panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Result: %v, %v\n", res.Status, string(data))
	respDataUnm := Response{}
	err = json.Unmarshal(data, &respDataUnm)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(respDataUnm)

	return respDataUnm.Result
}

func (b *BotClient) makeRequest(methodName string, body string) (*http.Response, error) {
	url := b.getURLWithoutMethod() + methodName
	resp, err := http.Post(url, "application/json", strings.NewReader(body))

	return resp, err
}

func (b *BotClient) getURLWithoutMethod() string {
	return b.config.TelegramURL + "bot" + b.config.AuthToken + "/"
}
