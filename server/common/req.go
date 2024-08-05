package common

import (
	"github.com/imroc/req/v3"
	"time"
)

func DefaultHttpReq(baseUrl string) *req.Client {
	return req.C().
		SetBaseURL(baseUrl).
		SetCommonRetryCount(3).
		SetCommonRetryBackoffInterval(100*time.Millisecond, 500*time.Millisecond)
}
