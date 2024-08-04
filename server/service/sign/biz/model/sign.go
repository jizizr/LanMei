package model

import "time"

type Sign struct {
	QQ           int64 `gorm:"primary_key"`
	Point        int64
	LastSignTime time.Time
}
