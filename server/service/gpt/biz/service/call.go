package service

import (
	"context"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/bot/biz/service"
	"github.com/jizizr/LanMei/server/service/gpt/biz/util"
)

const (
	costPerMillionTokens = 0.150
	costPerToken         = costPerMillionTokens / 1000000
)

func init() {
	go util.GPTQueue.Poll()
}

type CallService struct {
	ctx context.Context
} // NewCallService new CallService
func NewCallService(ctx context.Context) *CallService {
	return &CallService{ctx: ctx}
}

// Run create note info
func (s *CallService) Run(message *bot.Message) (resp bool, err error) {
	// Finish your business logic.
	resp = true
	var m *bot.MessageData
	for _, m = range message.Message {
		if m.Type == "at" {
			break
		}
	}
	if !(m != nil &&
		m.Type == "at" &&
		service.ParseStringToInt(*m.Data.Qq) == message.SelfId) {
		return
	}
	util.GPTQueue.Push(message)
	return
}
