package base

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//Model base model
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//User user model
type User struct {
	gorm.Model
	Email    string `sql:"not null;index" json:"email"`
	UID      string `sql:"not null;unique_index;type:char(36)" json:"uid"`
	Home     string `sql:"not null" json:"home"`
	Logo     string `sql:"not null" json:"logo"`
	Name     string `sql:"not null" json:"name"`
	Password string `sql:"not null;default:'-'" json:"-"`

	ProviderType string `sql:"not null;default:'unknown';index"`
	ProviderID   string `sql:"not null;index"`

	LastSignIn  time.Time  `sql:"not null" json:"last_sign_in"`
	SignInCount uint       `sql:"not null;default:0" json:"sign_in_count"`
	ConfirmedAt *time.Time `json:"confirmed_at"`
	LockedAt    *time.Time `json:"locked_at"`

	Permissions []Permission `json:"permissions"`
	Logs        []Log        `json:"logs"`
}

//IsConfirmed is confirmed?
func (p *User) IsConfirmed() bool {
	return p.ConfirmedAt != nil
}

//IsLocked is locked?
func (p *User) IsLocked() bool {
	return p.LockedAt != nil
}

//SetGravatar set gravatar logo
func (p *User) SetGravatar() {
	buf := md5.Sum([]byte(strings.ToLower(p.Email)))
	p.Logo = fmt.Sprintf("https://gravatar.com/avatar/%s.png", hex.EncodeToString(buf[:]))
}

func (p User) String() string {
	return fmt.Sprintf("%s<%s>", p.Name, p.Email)
}

//Log log model
type Log struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `sql:"not null" json:"-"`
	User      User      `json:"-"`
	Message   string    `sql:"not null" json:"message"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp" json:"created_at"`
}

//Role role model
type Role struct {
	Model

	Name         string `sql:"not null;index"`
	ResourceType string `sql:"not null;default:'-';index"`
	ResourceID   uint   `sql:"not null;default:0"`
}

func (p Role) String() string {
	return fmt.Sprintf("%s@%s://%d", p.Name, p.ResourceType, p.ResourceID)
}

//Permission permission model
type Permission struct {
	Model
	User   User
	UserID uint `sql:"not null"`
	Role   Role
	RoleID uint      `sql:"not null"`
	Begin  time.Time `sql:"not null;default:current_date;type:date"`
	End    time.Time `sql:"not null;default:'1000-1-1';type:date"`
}

//EndS endtime
func (p *Permission) EndS() string {
	return p.End.Format("2006-01-02")
}

//BeginS begintime
func (p *Permission) BeginS() string {
	return p.Begin.Format("2006-01-02")
}

//Enable is enable?
func (p *Permission) Enable() bool {
	now := time.Now()
	return now.After(p.Begin) && now.Before(p.End)
}

//Notice notice model
type Notice struct {
	Model
	Lang    string `sql:"not null;type:varchar(8);index" json:"lang"`
	Content string `sql:"not null;type:text" json:"content"`
}

//Setting model
type Setting struct {
	gorm.Model
	Key  string `sql:"not null;unique_index"`
	Val  []byte `sql:"not null"`
	Flag bool   `sql:"not null;default:false"`
}

//Locale model
type Locale struct {
	gorm.Model
	Lang    string `sql:"not null;type:varchar(8);index"`
	Code    string `sql:"not null;index"`
	Message string `sql:"not null;type:varchar(800)"`
}
