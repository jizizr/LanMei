package model

type Word struct {
	GroupID int64  `gorm:"primary_key;index"`
	Word    string `gorm:"primary_key"`
	Count   int
}
