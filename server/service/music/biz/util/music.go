package util

import (
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/imroc/req/v3"
	"github.com/jizizr/LanMei/server/common"
	"github.com/jizizr/LanMei/server/service/music/biz/model"
	"github.com/jizizr/LanMei/server/service/music/conf"
	"io"
	"time"
)

func init() {
	InitCookie()
}

var musicUrl = conf.GetConf().Music.Url

var client = common.DefaultHttpReq(musicUrl)
var musicStreamClient = req.C().
	R().
	SetRetryCount(3).
	SetRetryBackoffInterval(100*time.Millisecond, 500*time.Millisecond)

func InitCookie() {
	cookie := model.Ck{Data: conf.GetConf().Music.Ck}
	var resp model.CkResp
	r, err := client.R().SetBody(cookie).SetSuccessResult(&resp).Post("/user/SetCookie")
	if err != nil {
		panic(err)
	}
	if !r.IsSuccessState() || resp.Result != 100 {
		panic("init cookie error")
	}
}

func GetMusicInfo(musicName string) (musicInfo model.MusicInfo, err error) {
	r, err := client.R().SetSuccessResult(&musicInfo).SetQueryString("pageSize=1&t=0").
		AddQueryParam("key", musicName).
		Get("/search")
	if err != nil {
		return
	}
	if !r.IsSuccessState() || musicInfo.Result != 100 {
		err = errors.New("request api error")
	}
	return
}

func GetMusicStreamUrl(musicInfo *model.MusicInfo) (musicStreamUrl string, err error) {
	var musicStream model.MusicStream
	r, err := client.R().SetSuccessResult(&musicStream).
		AddQueryParam("id", musicInfo.Data.Lists[0].SongMid).
		AddQueryParam("mediaId", musicInfo.Data.Lists[0].MediaMid).
		Get("/song/url")
	if err != nil {
		return
	}
	if !r.IsSuccessState() || musicStream.Result != 100 {
		err = errors.New("request api error")
		return
	}
	musicStreamUrl = musicStream.Data
	return
}

func DownloadMusic(musicBuf io.Writer, musicStreamUrl string) (err error) {
	r, err := musicStreamClient.SetOutput(musicBuf).Get(musicStreamUrl)
	if err != nil {
		return
	}
	if !r.IsSuccessState() {
		klog.Error("send message error ", r.String())
		err = errors.New("request api error")
	}
	return
}
