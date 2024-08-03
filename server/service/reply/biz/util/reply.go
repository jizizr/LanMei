package util

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/LanMei/server/service/reply/biz/model"
	"github.com/jizizr/LanMei/server/service/reply/conf"
	"regexp"
	"strings"
)

type ReplyTable []ReplyRow

type ReplyRow interface {
	Match(words string) bool
	Reply() string
}

func NewReplyTable() *ReplyTable {
	r := make(ReplyTable, 0)
	return &r
}

func (r *ReplyTable) Match(words string) []string {
	var replies []string
	for _, row := range *r {
		if row.Match(words) {
			replies = append(replies, row.Reply())
		}
	}
	return replies
}

type RegexRow struct {
	matchPattern *regexp.Regexp
	reply        string
}

func NewRegexRow(pattern string, reply string) (*RegexRow, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RegexRow{
		matchPattern: reg,
		reply:        reply,
	}, nil
}

func (r *RegexRow) Match(words string) bool {
	return r.matchPattern.MatchString(words)
}

func (r *RegexRow) Reply() string {
	return r.reply
}

type MatchRow struct {
	matchWords string
	reply      string
}

func (m *MatchRow) Match(words string) bool {
	return m.matchWords == words
}

func (m *MatchRow) Reply() string {
	return m.reply
}

func NewMatchRow(words string, reply string) *MatchRow {
	return &MatchRow{
		matchWords: words,
		reply:      reply,
	}
}

type ContainRow struct {
	containWords string
	reply        string
}

func (c *ContainRow) Match(words string) bool {
	return strings.Contains(words, c.containWords)
}

func (c *ContainRow) Reply() string {
	return c.reply
}

func NewContainRow(words string, reply string) *ContainRow {
	return &ContainRow{
		containWords: words,
		reply:        reply,
	}
}

func UpdateReplyTable(replyTable *ReplyTable, reversion *int, token string) {
	var tableMetaInfo model.TableMetaInfo
	r, err := client.SetSuccessResult(&tableMetaInfo).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetPathParam("spreadsheetToken", conf.GetConf().TableAppInfo.SpreadsheetToken).
		Get("/sheets/v2/spreadsheets/{spreadsheetToken}/metainfo")
	if err != nil {
		klog.Error("get table meta info error ", err)
		return
	}
	if !r.IsSuccessState() || tableMetaInfo.Code != 0 {
		klog.Error("get table meta info error ", tableMetaInfo.Msg)
		return
	}
	if tableMetaInfo.Data.Properties.Revision == *reversion {
		return
	}
	var rawReplyTable model.ReplyTable
	sheetID := tableMetaInfo.Data.Sheets[0].SheetId
	r, err = client.SetSuccessResult(&rawReplyTable).
		SetPathParam("spreadsheetToken", conf.GetConf().TableAppInfo.SpreadsheetToken).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		AddQueryParam("ranges", sheetID).
		Get("/sheets/v2/spreadsheets/{spreadsheetToken}/values_batch_get")
	if err != nil {
		klog.Error("get reply table error ", err)
		return
	}
	if !r.IsSuccessState() || rawReplyTable.Code != 0 {
		klog.Error("get reply table error ", rawReplyTable.Msg)
		return
	}
	newReplyTable := make(ReplyTable, 0, len(rawReplyTable.Data.ValueRanges[0].Values)-1)
	for i, value := range rawReplyTable.Data.ValueRanges[0].Values[1:] {
		if len(value) < 3 {
			continue
		}

		switch value[2] {
		case "全字匹配":
			newReplyTable = append(newReplyTable, NewMatchRow(value[0], value[1]))
		case "包含文字":
			newReplyTable = append(newReplyTable, NewContainRow(value[0], value[1]))
		case "正则表达式":
			r, err := NewRegexRow(value[0], value[1])
			if err != nil {
				MarkInvalidRegexRow(sheetID, fmt.Sprintf("A%d", i+2), token)
				continue
			}
			newReplyTable = append(newReplyTable, r)
		default:
			newReplyTable = append(newReplyTable, NewMatchRow(value[0], value[1]))
		}
	}
	*reversion = rawReplyTable.Data.Revision
	*replyTable = newReplyTable
	return
}

func MarkInvalidRegexRow(sheetID string, ranges string, token string) {
	style := model.TableStyle{AppendStyle: model.AppendStyle{
		Range: fmt.Sprintf("%s!%s:%s", sheetID, ranges, ranges),
		Style: model.Style{
			BackColor: "#21d11f",
		},
	}}
	var resp model.StyleResp
	r, err := client.SetPathParam("spreadsheetToken", conf.GetConf().TableAppInfo.SpreadsheetToken).
		SetSuccessResult(&resp).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetBody(style).
		Put("/sheets/v2/spreadsheets/{spreadsheetToken}/style")
	if err != nil {
		klog.Error("mark invalid regex row error ", err)
	}
	if !r.IsSuccessState() || resp.Code != 0 {
		klog.Error("mark invalid regex row error ", resp.Msg)
	}
}
