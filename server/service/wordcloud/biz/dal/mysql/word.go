package mysql

import (
	"github.com/jizizr/LanMei/server/service/wordcloud/biz/model"
)

func init() {
	Init()
}

func GetWord(groupID int64) ([]model.Word, error) {
	var words []model.Word
	err := DB.Where("group_id = ?", groupID).Find(&words).Error
	return words, err
}
