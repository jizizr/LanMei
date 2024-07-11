package util

import (
	"errors"
	"fmt"
	"github.com/jizizr/LanMei/server/common"
	"regexp"
	"strings"
)

const BaseUrl = "https://hao.360.cn/histoday"

var client = common.DefaultHttpReq(BaseUrl)
var hisMatch = regexp.MustCompile(`</em>\.(.*?)</dt>`)

func GetHistory() (text string, err error) {
	resp, err := client.Get("/")
	if err != nil {
		return
	}
	if !resp.IsSuccessState() {
		err = errors.New("request api error")
		return
	}
	his := hisMatch.FindAllStringSubmatch(resp.String(), -1)
	lines := make([]string, len(his))
	for i := range his {
		lines[i] = fmt.Sprintf("%d. %s", i+1, his[i][1])
	}
	text = fmt.Sprintf("历史上的今天\n%s", strings.Join(lines, "\n"))
	return
}
