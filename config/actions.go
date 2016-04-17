package config

import (
	"crypto/aes"
	"fmt"
	"log"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/garyburd/redigo/redis"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/jinzhu/gorm"
	"github.com/jrallison/go-workers"
)

//IocAction ioc action
func IocAction(fn func(*martini.ClassicMartini, *cli.Context) error) func(*cli.Context) {
	return Action(func(cfg *Model, ctx *cli.Context) error {
		workers.Configure(map[string]string{
			"server":   fmt.Sprintf(cfg.Redis.URL()),
			"database": strconv.Itoa(cfg.Redis.Db),
			"pool":     strconv.Itoa(cfg.Workers.Pool),
			"process":  cfg.Workers.ID,
		})

		mux := martini.Classic()

		db, err := cfg.Database.Open()
		if err != nil {
			return err
		}

		db.LogMode(!cfg.IsProduction())

		mux.Map(db)
		mux.Map(cfg)
		mux.Map(cfg.Redis.Open())

		ak, err := cfg.Key(50, 32)
		if err != nil {
			return err
		}
		cip, err := aes.NewCipher(ak)
		if err != nil {
			return err
		}
		mux.Map(&web.Encryptor{Cip: cip})

		if err := web.Loop(func(en web.Engine) error {
			hd := en.Map(mux.Injector)
			if _, err := mux.Invoke(hd); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
		return fn(mux, ctx)
	})
}

//ENV 运行模式
var ENV = cli.StringFlag{
	Name:   "environment, e",
	Value:  "development",
	Usage:  "Specifies the environment to run this server under (test/development/production).",
	EnvVar: "ENV",
}

//EnvAction by env arg
func EnvAction(fn func(string, *cli.Context) error) func(*cli.Context) {
	return func(ctx *cli.Context) {
		log.Println("Begin...")
		if err := fn(ctx.String("environment"), ctx); err == nil {
			log.Println("Done!!!")
		} else {
			log.Fatalln(err)
		}
	}
}

//InvokeAction ivonke action
func InvokeAction(hd martini.Handler) func(*cli.Context) {
	return IocAction(func(mux *martini.ClassicMartini, ctx *cli.Context) error {
		rst, err := mux.Invoke(hd)
		if err == nil {
			return nil
		}
		val := rst[0].Interface()
		if val != nil {
			return val.(error)
		}
		return nil
	})
}

//Action config action
func Action(fn func(*Model, *cli.Context) error) func(*cli.Context) {
	return EnvAction(func(env string, ctx *cli.Context) error {
		var cfg Model
		if err := web.Load(fmt.Sprintf("%s.toml", env), &cfg); err != nil {
			return err
		}
		cfg.Env = env
		return fn(&cfg, ctx)
	})
}

//DbAction database action
func DbAction(fn func(*gorm.DB, *cli.Context) error) func(*cli.Context) {
	return Action(func(cfg *Model, ctx *cli.Context) error {
		db, err := cfg.Database.Open()
		if !cfg.IsProduction() {
			db.LogMode(true)
		}
		if err != nil {
			return err
		}
		return fn(db, ctx)
	})
}

//RedisAction redis action
func RedisAction(fn func(*redis.Pool, *cli.Context) error) func(*cli.Context) {
	return Action(func(cfg *Model, ctx *cli.Context) error {
		re := cfg.Redis.Open()
		return fn(re, ctx)
	})
}
