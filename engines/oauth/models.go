package oauth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/itpkg/web"
	"github.com/jinzhu/gorm"
)

//User user model
type User struct {
	gorm.Model
	Email    string `gorm:"not null;index;type:VARCHAR(255)" json:"email"`
	UID      string `gorm:"not null;unique_index;type:char(36)" json:"uid"`
	Home     string `gorm:"not null;type:VARCHAR(255)" json:"home"`
	Logo     string `gorm:"not null;type:VARCHAR(255)" json:"logo"`
	Name     string `gorm:"not null;type:VARCHAR(255)" json:"name"`
	Password string `gorm:"not null;default:'-';type:VARCHAR(500)" json:"-"`

	ProviderType string `gorm:"not null;default:'unknown';index;type:VARCHAR(255)"`
	ProviderID   string `gorm:"not null;index;type:VARCHAR(255)"`

	LastSignIn  *time.Time `json:"last_sign_in"`
	SignInCount uint       `gorm:"not null;default:0" json:"sign_in_count"`
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

//IsAvailable is available?
func (p *User) IsAvailable() bool {
	return p.IsConfirmed() && !p.IsLocked()
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
	UserID    uint      `gorm:"not null" json:"-"`
	User      User      `json:"-"`
	Message   string    `gorm:"not null;type:VARCHAR(255)" json:"message"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp" json:"created_at"`
}

//Role role model
type Role struct {
	web.Model

	Name         string `gorm:"not null;index;type:VARCHAR(255)"`
	ResourceType string `gorm:"not null;default:'-';index;type:VARCHAR(255)"`
	ResourceID   uint   `gorm:"not null;default:0"`
}

func (p Role) String() string {
	return fmt.Sprintf("%s@%s://%d", p.Name, p.ResourceType, p.ResourceID)
}

//Permission permission model
type Permission struct {
	web.Model
	User   User
	UserID uint `gorm:"not null"`
	Role   Role
	RoleID uint      `gorm:"not null"`
	Begin  time.Time `gorm:"not null;default:current_date;type:date"`
	End    time.Time `gorm:"not null;default:'1000-1-1';type:date"`
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
