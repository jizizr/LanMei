package mysql

import (
	"github.com/jizizr/LanMei/server/service/cut/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func init() {
	Init()
	NewWord()
}

func NewWord() {
	m := DB.Migrator()
	if m.HasTable(&model.Word{}) {
		return
	}
	err := m.CreateTable(&model.Word{})
	if err != nil {
		panic(err)
	}
}

func AddWords(groupID int64, words map[string]int) error {
	if len(words) == 0 {
		return nil
	}
	var wordModels []model.Word
	for word, count := range words {
		wordModels = append(wordModels, model.Word{GroupID: groupID, Word: word, Count: count})
	}

	return DB.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "group_id"}, {Name: "word"}}, // 冲突列
			DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("count + VALUES(count)")}),
		},
	).Create(&wordModels).Error
}
