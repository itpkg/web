package cms

import (
	"github.com/codegangsta/cli"
	"github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/jinzhu/gorm"
)

//Engine cms engine
type Engine struct {
}

//Map map objects
func (p *Engine) Map(inj inject.Injector) martini.Handler {
	return nil
}

//Mount mount to web
func (p *Engine) Mount(martini.Router) {

}

//Migrate call by db:migrate
func (p *Engine) Migrate(*gorm.DB) martini.Handler {
	return func(db *gorm.DB) {}
}

//Seed call by db:seed
func (p *Engine) Seed() martini.Handler {
	return func() {}
}

//Worker call by worker
func (p *Engine) Worker() martini.Handler {
	return func() {}
}

//Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	web.Register(&Engine{})
}
