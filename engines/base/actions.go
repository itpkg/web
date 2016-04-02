package base

import (
	"crypto/aes"
	"fmt"
	"log"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/jrallison/go-workers"
)

//IocAction ioc action
func IocAction(fn func(*martini.ClassicMartini, *cli.Context) error) func(*cli.Context) {
	return ConfigAction(func(cfg *Config, ctx *cli.Context) error {
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

//ConfigAction config action
func ConfigAction(fn func(*Config, *cli.Context) error) func(*cli.Context) {
	return EnvAction(func(env string, ctx *cli.Context) error {
		var cfg Config
		if err := web.Load(fmt.Sprintf("%s.toml", env), &cfg); err != nil {
			return err
		}
		cfg.Env = env
		return fn(&cfg, ctx)
	})
}
