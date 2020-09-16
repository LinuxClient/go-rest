package rest

import (
	"net/http"
	"time"
)

type Request struct {
	URL        string
	Method     string
	Token      string
	AuthScheme string
	Header     http.Header
	Time       time.Time
	Body       interface{}
	Result     interface{}
	Error      interface{}
	RawRequest *http.Request
	Cookies    []*http.Cookie

	Attempts int
}
