package base

import (
	"github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/itpkg/web/cache"
	"github.com/jinzhu/gorm"
)

//Engine base engine
type Engine struct {
}

//Map map objects
func (p *Engine) Map(inj inject.Injector) martini.Handler {
	return func(cfg *Config) {
		inj.Map(&cache.RedisProvider{
			Redis:  cfg.Redis.Open(),
			Prefix: "cache://",
		})
	}
}

//Mount mount to web
func (p *Engine) Mount(martini.Router) {

}

//Migrate call by db:migrate
func (p *Engine) Migrate() martini.Handler {
	return func(db *gorm.DB) {
		db.AutoMigrate(
			&Locale{}, &Setting{}, &Notice{},
			&User{}, &Role{}, &Permission{}, &Log{},
		)
		db.Model(&Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
		db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
		db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type_id", "name", "resource_type", "resource_id")
		db.Model(&Permission{}).AddUniqueIndex("idx_permissions_user_role", "user_id", "role_id")
	}
}

//Seed call by db:seed
func (p *Engine) Seed() martini.Handler {
	return func() {}
}

//Worker call by worker
func (p *Engine) Worker() martini.Handler {
	return func() {}
}

func init() {
	web.Register(&Engine{})
}
