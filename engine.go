package web

import (
	"github.com/codegangsta/cli"
	"github.com/codegangsta/inject"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
)

//Engine web module
type Engine interface {
	Map(inj inject.Injector) martini.Handler
	Mount(martini.Router)
	Migrate(*gorm.DB) martini.Handler
	Seed() martini.Handler
	Worker() martini.Handler
	Shell() []cli.Command
}

var engines []Engine

//Register registe engine
func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

//Loop loop for engines
func Loop(fn func(en Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
