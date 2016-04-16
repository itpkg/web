package dict

import (
	"github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/jinzhu/gorm"
)

//Engine dict engine
type Engine struct {
}

//Map map objects
func (p *Engine) Map(inj inject.Injector) martini.Handler {
	return func() {
		var dic Provider
		dic = &StarDict{Dir: "tmp/dic"}
		inj.Map(dic)
	}
}

//Migrate call by db:migrate
func (p *Engine) Migrate() martini.Handler {
	return func(db *gorm.DB) {

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
