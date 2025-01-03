package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/music/biz/util"
	"strings"
	"time"
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
		klog.Error(err)
		return
	}
	musicStreamUrl, err := util.GetMusicStreamUrl(&musicInfo)
	if err != nil {
		klog.Error(err)
		return
	}
	var MusicBuf bytes.Buffer
	err = util.DownloadMusic(&MusicBuf, musicStreamUrl)
	if err != nil {
		return
	}
	msg := common.NewMsg(message)
	msg.Message = fmt.Sprintf(
		"[CQ:file,file=base64://%s,name=%s.mp3]",
		base64.URLEncoding.EncodeToString(MusicBuf.Bytes()),
		musicInfo.Data.Lists[0].SongName,
	)
	msgID, err := msg.SendMessage()
	if err != nil {
		return
	}
	msg = common.NewMsg(message)
	msg.Message = fmt.Sprintf(
		"歌手：%s\n发布时间：%s",
		musicInfo.Data.Lists[0].Singers[0].Name,
		time.Unix(musicInfo.Data.Lists[0].PubTime, 0).Format("2006-01-02"),
	)
	msg.Reply(msgID).SendMessage()
	return
}
