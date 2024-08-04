package common

import (
	"fmt"
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
	MessageType string `json:"message_type,omitempty"`
	MessageID   int64  `json:"message_id,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	GroupID     *int64 `json:"group_id,omitempty"`
	Message     string `json:"message,omitempty"`
	Duration    uint32 `json:"duration,omitempty"`
	AutoEscape  bool   `json:"auto_escape,omitempty"`
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

func NewMsg(message *bot.Message) *Msg {
	return &Msg{
		MessageType: message.MessageType,
		MessageID:   message.MessageId,
		UserID:      message.UserId,
		GroupID:     message.GroupId,
		Message:     "",
		AutoEscape:  false,
	}
}

func (m *Msg) At(uid ...int64) *Msg {
	var id int64
	if len(uid) > 0 {
		id = uid[0]
	} else {
		id = m.UserID
	}
	m.Message = fmt.Sprintf("[CQ:at,qq=%d] %s", id, m.Message)
	return m
}

func (m *Msg) Reply(mid ...int64) *Msg {
	var id int64
	if len(mid) > 0 {
		id = mid[0]
	} else {
		id = m.MessageID
	}
	m.Message = fmt.Sprintf("[CQ:reply,id=%d]%s", id, m.Message)
	return m
}

func (m *Msg) send(path string) (int64, error) {
	content, err := sonic.Marshal(m)
	if err != nil {
		klog.Error("marshal message error ", err, m)
		return 0, err
	}
	var resp Response
	r, err := msgClient.SetBodyBytes(content).SetSuccessResult(&resp).Post(path)
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

func (m *Msg) SendMessage() (int64, error) {
	return m.send("/send_msg")
}

func (m *Msg) SendBan(d uint32) (int64, error) {
	m.Duration = d
	return m.send("/set_group_ban")
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
