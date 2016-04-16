package dict

import "github.com/itpkg/web/engines/base"

//Note word note
type Note struct {
	base.Model
	Title  string `gorm:"not null;index;type:VARCHAR(255)" json:"title"`
	Body   string `gorm:"not null;type:TEXT" json:"body"`
	UserID uint   `gorm:"not null" json:"userId"`
	User   base.User
}

//TableName table name
func (Note) TableName() string {
	return "dict_notes"
}
