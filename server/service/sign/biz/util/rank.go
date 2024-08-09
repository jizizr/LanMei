package util

import (
	"fmt"
	"github.com/jizizr/LanMei/server/common"
	"github.com/jizizr/LanMei/server/rpc_gen/kitex_gen/bot"
	"github.com/jizizr/LanMei/server/service/sign/biz/dal/mysql"
	"github.com/jizizr/LanMei/server/service/sign/biz/model"
	"strconv"
	"strings"
)

func GetGroupMemberNickName(groupID, userID int64) string {
	var member model.Member
	r, err := common.MsgClient.R().SetFormData(
		map[string]string{
			"group_id": strconv.FormatInt(groupID, 10),
			"user_id":  strconv.FormatInt(userID, 10),
		},
	).
		SetSuccessResult(&member).
		Post("/get_group_member_info")
	if err != nil || !r.IsSuccessState() || member.Retcode != 0 || member.Status != "ok" {
		return ""
	}
	if member.Data.Card != "" {
		return member.Data.Card
	}
	return member.Data.Nickname
}

func GetRankN(message *bot.Message, number int) (text string, err error) {
	rankList, err := mysql.GetRankTopN(number)
	if err != nil {
		return
	}
	rankNameList := make([]string, 0, len(rankList))
	for i, rank := range rankList {
		rankNameList = append(rankNameList, fmt.Sprintf("第%d名：%s，总积分：%d", i+1, GetGroupMemberNickName(*message.GroupId, rank.QQ), rank.Point))
	}
	text = fmt.Sprintf("群内积分排行榜前%d名：\n%s", len(rankNameList), strings.Join(rankNameList, "\n"))
	return
}
