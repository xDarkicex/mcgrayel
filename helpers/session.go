package helpers

import (
	"math/rand"
	"time"

	"github.com/gorilla/sessions"
)

var (
	_store *sessions.CookieStore
)

// Store returns a cookie
func Store() *sessions.CookieStore {
	if _store == nil {
		_store = sessions.NewCookieStore([]byte(RandStringRunes(300)))
	}
	return _store
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
