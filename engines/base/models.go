package base

import "github.com/itpkg/web"

//Notice notice model
type Notice struct {
	web.Model
	Lang    string `gorm:"not null;type:varchar(8);index" json:"lang"`
	Content string `gorm:"not null;type:text" json:"content"`
}
