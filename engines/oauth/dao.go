package oauth

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Dao database operations
type Dao struct {
	Db *gorm.DB
}

//Is is?
func (p *Dao) Is(user uint, name string) bool {
	return p.Can(user, name, "-", 0)
}

//Can can?
func (p *Dao) Can(user uint, name string, rty string, rid uint) bool {
	var r Role
	if p.Db.Where("name = ? AND resource_type = ? AND resource_id = ?", name, rty, rid).First(&r).RecordNotFound() {
		return false
	}
	var pm Permission
	if p.Db.Where("user_id = ? AND role_id = ?", user, r.ID).First(&pm).RecordNotFound() {
		return false
	}

	return pm.Enable()
}

//Role insure role name exists
func (p *Dao) Role(name string, rty string, rid uint) (*Role, error) {
	var e error
	r := Role{}
	db := p.Db
	if db.Where("name = ? AND resource_type = ? AND resource_id = ?", name, rty, rid).First(&r).RecordNotFound() {
		r = Role{
			Name:         name,
			ResourceType: rty,
			ResourceID:   rid,
		}
		e = db.Create(&r).Error

	}
	return &r, e
}

//Deny set deny permission
func (p *Dao) Deny(role uint, user uint) error {
	return p.Db.Where("role_id = ? AND user_id = ?", role, user).Delete(Permission{}).Error
}

//Allow set allow permission
func (p *Dao) Allow(role uint, user uint, dur time.Duration) error {
	begin := time.Now()
	end := begin.Add(dur)
	var count int
	p.Db.Model(&Permission{}).Where("role_id = ? AND user_id = ?", role, user).Count(&count)
	if count == 0 {
		return p.Db.Create(&Permission{
			UserID: user,
			RoleID: role,
			Begin:  begin,
			End:    end,
		}).Error
	}
	return p.Db.Model(&Permission{}).Where("role_id = ? AND user_id = ?", role, user).UpdateColumns(map[string]interface{}{"begin": begin, "end": end}).Error

}
