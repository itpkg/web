package oauth

import (
	"log"

	"github.com/codegangsta/inject"
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/itpkg/web/engines/base"
	"github.com/itpkg/web/token"
	"github.com/jinzhu/gorm"
)

//Engine oauth engine
type Engine struct {
}

//Map map objects
func (p *Engine) Map(inj inject.Injector) martini.Handler {
	return func(re *redis.Pool, cfg *base.Config, lg *log.Logger, db *gorm.DB) {
		key, err := cfg.Key(60, 17)
		if err != nil {
			lg.Fatal(err)
		}
		inj.Map(&token.Jwt{
			Provider: &token.RedisProvider{Redis: re},
			Key:      key,
		})

		inj.Map(&Dao{Db: db})
	}
}

//Migrate call by db:migrate
func (p *Engine) Migrate() martini.Handler {
	return func(db *gorm.DB) {
		db.AutoMigrate(
			&User{}, &Role{}, &Permission{}, &Log{},
		)
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
