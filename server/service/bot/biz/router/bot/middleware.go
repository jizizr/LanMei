// Code generated by hertz generator.

package bot

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jizizr/LanMei/server/service/bot/biz/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.SigAuth(),
	}
}

func _getmessageMw() []app.HandlerFunc {
	// your code...
	return nil
}
