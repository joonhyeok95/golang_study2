package domain

import (
	"time"
)

type TMemberTemp struct {
	// gorm.Model           // 추적에 용이함
	// ID         uint      `gorm:"primaryKey"`
	UserId     string    `gorm:"column:USER_ID;primaryKey"`
	FirstName  string    `gorm:"column:FIRST_NAME"`
	LastName   string    `gorm:"column:LAST_NAME"`
	Email      string    `gorm:"column:EMAIL"`
	CustomDate time.Time `gorm:"column:CUSTOM_DATE"`
}

func (TMemberTemp) TableName() string {
	return "t_member_temp"
}
