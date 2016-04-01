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
		mux.Map(&web.Aes{Cip: cip})
		mux.Map(&web.BytesSerial{})

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

//Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init config file",
			Flags:   []cli.Flag{ENV},
			Action: EnvAction(func(env string, _ *cli.Context) error {
				sec, err := web.Random(512)
				if err != nil {
					return err
				}
				return web.Store(fmt.Sprintf("%s.toml", env), &Config{
					Secrets: web.ToBase64(sec),
					HTTP: HTTP{
						Host: "localhost",
						Port: 8080,
					},
					Database: Database{
						Type: "postgres",
						Args: map[string]string{
							"dbname":  "itpkg_dev",
							"sslmode": "disable",
							"user":    "postgres",
						},
					},
					Redis: Redis{
						Host: "localhost",
						Port: 6379,
						Db:   2,
					},
					ElasticSearch: ElasticSearch{
						Host:  "localhost",
						Port:  9200,
						Index: "itpkg-dev",
					},
					Workers: Workers{
						ID:     "itpkg-workers",
						Pool:   15,
						Queues: map[string]int{"default": 1, "emails": 2},
					},
				})
			}),
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the web server",
			Flags:   []cli.Flag{ENV},
			Action: EnvAction(func(env string, _ *cli.Context) error {
				return nil
			}),
		},
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Aliases: []string{"n"},
					Usage:   "create database",
					Flags:   []cli.Flag{ENV},
					Action: ConfigAction(func(cfg *Config, ctx *cli.Context) error {
						switch cfg.Database.Type {
						case "postgres":
							c, a := cfg.Database.Execute(fmt.Sprintf("CREATE DATABASE %s WITH ENCODING='UTF8'", cfg.Database.Args["dbname"]))
							return web.Shell(c, a...)
						default:
							return fmt.Errorf("bad database type %s", cfg.Database.Type)
						}
					}),
				},
				{
					Name:    "console",
					Aliases: []string{"c"},
					Usage:   "start a console for the database",
					Flags:   []cli.Flag{ENV},
					Action: ConfigAction(func(cfg *Config, ctx *cli.Context) error {
						c, a := cfg.Database.Console()
						return web.Shell(c, a...)
					}),
				},
				{
					Name:    "drop",
					Aliases: []string{"d"},
					Usage:   "drop database",
					Flags:   []cli.Flag{ENV},
					Action: ConfigAction(func(cfg *Config, ctx *cli.Context) error {
						switch cfg.Database.Type {
						case "postgres":
							c, a := cfg.Database.Execute(fmt.Sprintf("DROP DATABASE %s", cfg.Database.Args["dbname"]))
							return web.Shell(c, a...)
						default:
							return fmt.Errorf("bad database type %s", cfg.Database.Type)
						}
					}),
				},
				{
					Name:    "migrate",
					Aliases: []string{"m"},
					Usage:   "migrate the database",
					Flags:   []cli.Flag{ENV},
					Action: IocAction(func(mux *martini.ClassicMartini, ctx *cli.Context) error {
						return web.Loop(func(en web.Engine) error {
							_, err := mux.Invoke(en.Migrate())
							return err
						})
					}),
				},
				{
					Name:    "seed",
					Aliases: []string{"s"},
					Usage:   "load the seed data",
					Flags:   []cli.Flag{ENV},
					Action: IocAction(func(mux *martini.ClassicMartini, ctx *cli.Context) error {
						return web.Loop(func(en web.Engine) error {
							_, err := mux.Invoke(en.Seed())
							return err
						})

					}),
				},
			},
		},
	}
}
