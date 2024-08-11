package utils

import (
	"github.com/jizizr/LanMei/server/service/bot/hertz_gen/bot"
	"slices"
	"strings"
)

func FixMessage(req *bot.Message) {
	for i, message := range req.Message {
		if message.Type == "at" {
			return
		}
		if message.Type != "text" {
			continue
		}
		texts := strings.SplitN(*message.Data.Text, " ", 2)
		if len(texts) == 0 {
			return
		}
		if strings.HasPrefix(strings.TrimSpace(texts[0]), "@") {
			MsgAtData := bot.NewMessageData()
			MsgAtData.Type = "at"
			MsgAtData.Data = bot.NewData()
			MsgAtData.Data.Text = &texts[0]
			qqZero := "0"
			MsgAtData.Data.Qq = &qqZero
			if len(texts) == 1 {
				req.Message = slices.Replace(req.Message, i, i+1, MsgAtData)
			} else {
				MsgTextData := bot.NewMessageData()
				MsgTextData.Type = "text"
				MsgTextData.Data = bot.NewData()
				MsgTextData.Data.Text = &texts[1]
				req.Message = slices.Replace(req.Message, i, i+1, MsgAtData, MsgTextData)
			}
		}
		return
	}
}
