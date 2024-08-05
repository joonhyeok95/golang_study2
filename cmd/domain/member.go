package domain

import (
	"time"
)

type TMember struct {
	// gorm.Model           // ID,CreateAt,UpdateAt,DeleteAt 자동 생성
	// ID         uint      `gorm:"primaryKey"`
	UserId     string    `gorm:"column:USER_ID;primaryKey"`
	FirstName  string    `gorm:"column:FIRST_NAME"`
	LastName   string    `gorm:"column:LAST_NAME"`
	Email      string    `gorm:"column:EMAIL"`
	CustomDate time.Time `gorm:"column:CUSTOM_DATE"`
}

func (TMember) TableName() string {
	return "t_member"
}
