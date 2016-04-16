package dict

import (
	"github.com/itpkg/web"
	"github.com/itpkg/web/engines/oauth"
)

//Note word note
type Note struct {
	web.Model
	Title  string `gorm:"not null;index;type:VARCHAR(255)" json:"title"`
	Body   string `gorm:"not null;type:TEXT" json:"body"`
	UserID uint   `gorm:"not null" json:"userId"`
	User   oauth.User
}

//TableName table name
func (Note) TableName() string {
	return "dict_notes"
}
