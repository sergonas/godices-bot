package godicesbot

var (
	userStorage map[string]string = make(map[string]string)
)

type UserMapping struct {
	ID           int
	TelegramName string
	Login        string
}

// LoginByTelegram search inner login by telegram username
func LoginByTelegram(telegram string) (string, bool) {
	val, ok := userStorage[telegram]
	return val, ok
}

// TelegramByLogin search telegram username by inner login
func TelegramByLogin(login string) (string, bool) {
	for key, value := range userStorage {
		if value == login {
			return key, true
		}
	}
	return "", false
}

// StoreLink stores link between telegram username and inner login
func StoreLink(telegram, login string) {
	userStorage[telegram] = login
}
