package rest

import (
	"net/http"
	"time"
)

type Response struct {
	Request     *Request
	RawResponse *http.Response
	body        []byte
	size        int64
	receivedAt  time.Time
}

func (response *Response) Body() []byte {
	if response.RawResponse == nil {
		return []byte{}
	}
	return response.body
}
