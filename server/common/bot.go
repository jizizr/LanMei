package common

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/imroc/req/v3"
	"time"
)

const (
	BaseUrl = "http://127.0.0.1:3000"
)

var msgClient = req.C().
	SetBaseURL(BaseUrl).
	R().
	SetRetryCount(3).
	SetRetryBackoffInterval(100*time.Millisecond, 500*time.Millisecond)

type Msg struct {
	MessageType string `json:"message_type"`
	UserID      int64  `json:"user_id"`
	GroupID     *int64 `json:"group_id"`
	Message     string `json:"message"`
	AutoEscape  bool   `json:"auto_escape"`
}

type Response struct {
	Status  string      `json:"status"`
	Retcode int         `json:"retcode"`
	Data    Data        `json:"data"`
	Message string      `json:"message"`
	Wording string      `json:"wording"`
	Echo    interface{} `json:"echo"`
}

type Data struct {
	MessageId int64 `json:"message_id"`
}

func (m *Msg) Send() (int64, error) {
	content, err := sonic.Marshal(m)
	if err != nil {
		klog.Error("marshal message error ", err, m)
		return 0, err
	}
	var resp Response
	r, err := msgClient.SetBodyBytes(content).SetSuccessResult(&resp).Post("/send_msg")
	if err != nil {
		klog.Error("send message error ", err, m)
		return 0, err
	}
	if !r.IsSuccessState() {
		klog.Error("send message error ", r.String())
		return 0, err
	}
	if resp.Data.MessageId == 0 {
		klog.Error("send message error ", resp.Message)
		return 0, err
	}
	return resp.Data.MessageId, nil
}
