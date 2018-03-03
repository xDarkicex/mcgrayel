package helpers

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

// Controller Struct
type Controller struct {
	Controller interface{}
}

// RouterArgs These are the arguments passed in from a router.
type RouterArgs struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
	Session  *sessions.Session
	User     interface{}
}

// Flash ...
type Flash struct {
	Type    string
	Message string
}

// RoutesHandler for handling padding multiple objects into routes
type RoutesHandler func(a RouterArgs)

//ContactForm struct for passing data to emails
type ContactForm struct {
	Department string
	Body       string
	Name       string
	Telephone  string
	Product    string
	Email      string
}

//Secret containts password
type Secret struct {
	Users []struct {
		Password string
	}
}

//RecaptchaResponse is for google captcha
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}
