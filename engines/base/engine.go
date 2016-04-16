package base

import (
	"log"

	"github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/itpkg/web/cache"
	"github.com/itpkg/web/i18n"
	"github.com/itpkg/web/settings"
	"github.com/jinzhu/gorm"
)

//Engine base engine
type Engine struct {
}

//Map map objects
func (p *Engine) Map(inj inject.Injector) martini.Handler {
	return func(cfg *Config, enc *web.Encryptor, db *gorm.DB, lg *log.Logger) {
		t := i18n.I18n{
			Provider: &i18n.DatabaseProvider{
				Db:     db,
				Logger: lg,
			},
			Locales: make(map[string]map[string]string),
			Logger:  lg,
		}
		t.Load("locales")
		inj.Map(&t)

		inj.Map(&settings.DatabaseProvider{
			Db:  db,
			Enc: enc,
		})

		inj.Map(&cache.RedisProvider{
			Redis:  cfg.Redis.Open(),
			Prefix: "cache://",
		})

		inj.Map(&Dao{Db: db})
	}
}

//Migrate call by db:migrate
func (p *Engine) Migrate() martini.Handler {
	return func(db *gorm.DB) {
		db.AutoMigrate(
			&i18n.Locale{}, &settings.Model{},
			&Notice{},
			&User{}, &Role{}, &Permission{}, &Log{},
		)
		db.Model(&i18n.Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
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
