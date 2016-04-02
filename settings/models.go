package settings

import (
	"github.com/jinzhu/gorm"
)

// Model Setting enyry
type Model struct {
	gorm.Model

	Key  string `gorm:"not null;unique"`
	Val  []byte `gorm:"not null"`
	Flag bool   `gorm:"not null"`
}

//TableName custom table name
func (Model) TableName() string {
	return "settings"
}
