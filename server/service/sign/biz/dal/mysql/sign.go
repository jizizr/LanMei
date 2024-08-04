package mysql

import (
	"errors"
	"github.com/jizizr/LanMei/server/service/sign/biz/model"
	"gorm.io/gorm"
	"time"
)

func init() {
	Init()
	NewSign()
}

func NewSign() {
	m := DB.Migrator()
	if m.HasTable(&model.Sign{}) {
		return
	}
	err := m.CreateTable(&model.Sign{})
	if err != nil {
		panic(err)
	}
}

func getUserSignRecord(tx *gorm.DB, userQQ int64) (*model.Sign, error) {
	var signRecord model.Sign
	err := tx.Set("gorm:query_option", "FOR UPDATE").Where("qq = ?", userQQ).First(&signRecord).Error
	if err != nil {
		return nil, err
	}
	return &signRecord, nil
}

func createNewSignRecord(tx *gorm.DB, userQQ int64, currentTime time.Time, point int64) error {
	newSignRecord := model.Sign{
		QQ:           userQQ,
		Point:        point,
		LastSignTime: currentTime,
	}
	return tx.Create(&newSignRecord).Error
}

func updateSignRecord(tx *gorm.DB, signRecord *model.Sign, currentTime time.Time, point int64) error {
	signRecord.Point += point
	signRecord.LastSignTime = currentTime
	return tx.Save(signRecord).Error
}

func isAlreadySignedToday(lastSignTime, currentTime time.Time) bool {
	year, month, day := currentTime.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, currentTime.Location())
	return !lastSignTime.Before(startOfDay)
}

func getUserRankByPoints(userPoints int64) (int64, error) {
	var rank int64

	// 查询比当前用户积分多的用户数量
	err := DB.Model(&model.Sign{}).
		Where("point > ?", userPoints).
		Count(&rank).Error

	if err != nil {
		return 0, err
	}

	// 用户的排名是比他积分多的用户数量 + 1
	rank = rank + 1

	return rank, nil
}

func SignIn(userQQ int64, point int64) (alreadySignedToday bool, pointNow, rank int64, err error) {
	currentTime := time.Now()
	var signRecord *model.Sign
	err = DB.Transaction(func(tx *gorm.DB) error {
		signRecord, err = getUserSignRecord(tx, userQQ)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err // 其他错误，事务回滚
			}
			// 用户没有签到记录，创建新记录
			if createErr := createNewSignRecord(tx, userQQ, currentTime, point); createErr != nil {
				return createErr
			}
			return nil
		}

		// 检查是否已经在今天签到过
		if isAlreadySignedToday(signRecord.LastSignTime, currentTime) {
			alreadySignedToday = true
			return nil
		}

		// 更新签到记录
		return updateSignRecord(tx, signRecord, currentTime, point)
	})
	if signRecord != nil {
		pointNow = signRecord.Point
	} else {
		pointNow = point
	}
	rank, err = getUserRankByPoints(pointNow)
	return
}
