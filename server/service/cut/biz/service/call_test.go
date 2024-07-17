package service

import (
	"context"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"testing"
)

func TestCall_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCallService(ctx)
	// init req and assert value
	text := "测试"
	var gid int64 = 1111111111
	message := &bot.Message{
		SelfId:      0,
		UserId:      999,
		Time:        0,
		MessageId:   0,
		MessageSeq:  0,
		RealId:      0,
		MessageType: "",
		Sender:      nil,
		RawMessage:  "",
		Font:        0,
		SubType:     nil,
		Message: []*bot.MessageData{
			&bot.MessageData{
				Type: "text",
				Data: &bot.Data{
					Text:  &text,
					Id:    nil,
					Qq:    nil,
					Name:  nil,
					Type:  nil,
					Url:   nil,
					Audio: nil,
					Title: nil,
				},
			},
		},
		MessageFormat: "",
		PostType:      "",
		GroupId:       &gid,
		NoticeType:    nil,
		OperatorId:    nil,
	}
	resp, err := s.Run(message)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
