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

var musicUrl = conf.GetConf().Music.Url

var client = common.DefaultHttpReq(musicUrl).SetQueryString("n=1&q=7")
var musicStreamClient = req.C().
	R().
	SetRetryCount(3).
	SetRetryBackoffInterval(100*time.Millisecond, 500*time.Millisecond)

func GetMusicInfo(musicName string) (musicInfo model.MusicInfo, err error) {
	r, err := client.SetSuccessResult(&musicInfo).AddQueryParam("word", musicName).Get("/")
	if err != nil {
		return
	}
	if !r.IsSuccessState() {
		klog.Error("send message error ", r.String())
		err = errors.New("request api error")
	}
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
