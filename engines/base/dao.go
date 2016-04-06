package base

import "github.com/jinzhu/gorm"

//Dao database operations
type Dao struct {
	Db *gorm.DB
}
