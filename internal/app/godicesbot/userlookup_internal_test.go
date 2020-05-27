package godicesbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserlookup_MainRoute(t *testing.T) {
	StoreLink("one", "another")
	telegram, _ := TelegramByLogin("another")
	assert.Equal(t, "one", telegram)
	login, _ := LoginByTelegram("one")
	assert.Equal(t, "another", login)
}
