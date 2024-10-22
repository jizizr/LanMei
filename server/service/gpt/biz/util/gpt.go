package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/common"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/gpt/biz/model"
	"github.com/jizizr/LanMei/server/service/gpt/conf"
	openai "github.com/sashabaranov/go-openai"
	"strings"
)

const (
	costPerMillionTokens = 0.150
	costPerToken         = costPerMillionTokens / 1000000
)

var client *openai.Client
var lastMessageID [2]string
var runReq = openai.RunRequest{
	AssistantID: conf.GetConf().GPT.AssistantID,
}
var GPTQueue = model.NewGPTQueue(HandlerMessage)

func init() {
	config := openai.DefaultConfig(conf.GetConf().GPT.Token)
	config.BaseURL = conf.GetConf().GPT.Url
	client = openai.NewClientWithConfig(config)
}

func AddMessage(msg *bot.Message) error {
	text := strings.TrimSpace(common.ExtractText(msg))
	if text == "" {
		return errors.New("empty message")
	}
	messageReq := openai.MessageRequest{
		Role:    "user",
		Content: text,
	}
	m, err := client.CreateMessage(
		context.Background(),
		conf.GetConf().GPT.ThreadID,
		messageReq,
	)
	if err != nil {
		return err
	} else {
		lastMessageID[0] = lastMessageID[1]
		lastMessageID[1] = m.ID
		return nil
	}
}

func RunThread() (openai.Run, error) {
	return client.CreateRun(context.Background(), conf.GetConf().GPT.ThreadID, runReq)
}

func RetrieveRun(run openai.Run) (openai.Run, error) {
	return client.RetrieveRun(context.Background(), conf.GetConf().GPT.ThreadID, run.ID)
}

func ListMessage(limiter int) (openai.MessagesList, error) {
	return client.ListMessage(context.Background(), conf.GetConf().GPT.ThreadID, &limiter, nil, nil, nil)
}

func HandlerMessage(msg *bot.Message) {
	err := AddMessage(msg)
	if err != nil {
		klog.Error(err)
		return
	}
	defer func() {
		if lastMessageID[0] == "" {
			return
		}
		go client.DeleteMessage(context.Background(), conf.GetConf().GPT.ThreadID, lastMessageID[0])
	}()
	run, err := RunThread()
	if err != nil {
		klog.Error(err)
		return
	}
	for run.Status == openai.RunStatusInProgress || run.Status == openai.RunStatusQueued {
		run, err = RetrieveRun(run)
		if err != nil {
			klog.Error(err)
			return
		}
	}
	if run.Status != openai.RunStatusCompleted {
		klog.Error("run status is not complete ", run.Status)
		hlog.Error(run)
		return
	}
	messages, err := ListMessage(1)
	if err != nil {
		klog.Error(err)
		return
	}
	go func() {
		if len(messages.Messages) == 0 {
			klog.Error("no message")
			return
		}
		if len(messages.Messages[0].Content) == 0 {
			klog.Error("no content")
			return
		}
		m := common.NewMsg(msg)
		m.Message = fmt.Sprintf("%s\n\n共消耗%dToken ≈ %.5f$",
			messages.Messages[0].Content[0].Text.Value,
			run.Usage.TotalTokens,
			float64(run.Usage.TotalTokens)*costPerToken,
		)
		m.Reply().SendMessage()
		client.DeleteMessage(context.Background(), conf.GetConf().GPT.ThreadID, messages.Messages[0].ID)
	}()
}
