package model

import (
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
)

type GPTQueue struct {
	messageChan chan *bot.Message
	handlerFunc func(*bot.Message)
}

func NewGPTQueue(f func(msg *bot.Message)) *GPTQueue {
	return &GPTQueue{
		messageChan: make(chan *bot.Message, 10),
		handlerFunc: f,
	}
}

func (q *GPTQueue) Push(msg *bot.Message) {
	q.messageChan <- msg
}

func (q *GPTQueue) Poll() {
	for {
		select {
		case msg := <-q.messageChan:
			q.handlerFunc(msg)
		}
	}
}
