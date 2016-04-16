package oauth

import (
	"log"

	"github.com/codegangsta/inject"
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/itpkg/web/engines/base"
	"github.com/itpkg/web/token"
)

//Engine base engine
type Engine struct {
}

//Map map objects
func (p *Engine) Map(inj inject.Injector) martini.Handler {
	return func(re *redis.Pool, cfg *base.Config, lg *log.Logger) {
		key, err := cfg.Key(60, 17)
		if err != nil {
			lg.Fatal(err)
		}
		inj.Map(&token.Jwt{
			Provider: &token.RedisProvider{Redis: re},
			Key:      key,
		})
	}
}

//Migrate call by db:migrate
func (p *Engine) Migrate() martini.Handler {
	return func() {
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
