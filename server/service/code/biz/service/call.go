package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/common"
	bot "github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/code/biz/util"
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
	msg := common.NewMsg(message)
	text := common.ExtractText(message)
	language, code := util.Parse(text)
	if _, ok := util.CodeType[language]; !ok || code == "" {
		msg.Message =
			`输入有误，目前仅支持py/php/java/cpp/js/c#/c/go/asm/ats/bash/clisp/clojure/cobol/coffeescript/crystal/d/elixir/elm/erlang/fsharp/groovy/guide/hare/haskell/idris/julia/kotlin/lua/mercury/nim/nix/ocaml/pascal/perl/raku/ruby/rust/sac/scala/swift/typescript/zig/plaintext
格式:
/code [语言]
[代码]
`
		msg.Reply().At().SendMessage()
		return
	}
	result, err := util.Run(code, language)
	if err != nil {
		klog.Error(err)
		return
	}
	msg.Message = result
	msg.AutoEscape = true
	msg.Reply().SendMessage()
	return
}
