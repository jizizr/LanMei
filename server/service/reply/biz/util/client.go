package util

import "github.com/jizizr/LanMei/server/common"

const BaseUrl = "https://open.feishu.cn/open-apis"

var client = common.DefaultHttpReq(BaseUrl).
	SetHeader("Content-Type", "application/json; charset=utf-8")
