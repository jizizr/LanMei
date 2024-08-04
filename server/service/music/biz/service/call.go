package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/music/biz/util"
	"strings"
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
	messageArr := strings.SplitN(common.ExtractText(message), " ", 2)
	if len(messageArr) < 2 {
		msg := common.NewMsg(message)
		msg.Message = "使用方法:\n/music [歌名]"
		msg.Reply().At().SendMessage()
		return
	}
	musicInfo, err := util.GetMusicInfo(strings.TrimSpace(messageArr[1]))
	if err != nil {
		return
	}
	var MusicBuf bytes.Buffer
	err = util.DownloadMusic(&MusicBuf, musicInfo.Data.Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := common.NewMsg(message)
	musicData := musicInfo.Data
	msg.Message = fmt.Sprintf(
		"[CQ:file,file=base64://%s,name=%s.mp3]",
		base64.URLEncoding.EncodeToString(MusicBuf.Bytes()),
		musicData.Song,
	)
	_, err = msg.Reply().SendMessage()
	if err != nil {
		return
	}
	return
}
