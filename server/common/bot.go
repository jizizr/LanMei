package common

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"strings"
)

const (
	BaseUrl = "http://127.0.0.1:3000"
)

var msgClient = DefaultHttpReq(BaseUrl)

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

func NewMsg(message *bot.Message) Msg {
	return Msg{
		MessageType: message.MessageType,
		UserID:      message.UserId,
		GroupID:     message.GroupId,
		Message:     "",
		AutoEscape:  false,
	}
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

func IsBot(user int64) bool {
	return user == 3328144510 || user >= 2854196301 && user <= 2854216399 || user == 66600000 || user >= 3889000000 && user <= 3889999999
}

func ExtractText(msg *bot.Message) string {
	texts := make([]string, 0)
	for _, m := range msg.Message {
		if m.Type != "text" {
			continue
		}
		texts = append(texts, *m.Data.Text)
	}
	return strings.Join(texts, "")
}
