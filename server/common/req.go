package common

import (
	"github.com/imroc/req/v3"
	"time"
)

func DefaultHttpReq(baseUrl string) *req.Request {
	return req.C().
		SetBaseURL(baseUrl).
		R().
		SetRetryCount(3).
		SetRetryBackoffInterval(100*time.Millisecond, 500*time.Millisecond)
}
