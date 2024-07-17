package service

import (
	"context"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/cut/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/cut/biz/util"
)

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
	if message.GroupId == nil {
		return
	}
	if common.IsBot(message.UserId) {
		return
	}
	words := util.Cut(common.ExtractText(message))
	err = mysql.AddWords(*message.GroupId, words)
	return
}
