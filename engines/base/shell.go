package base

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/attilaolah/strict"
	"github.com/codegangsta/cli"
	"github.com/go-martini/martini"
	"github.com/itpkg/web"
	"github.com/itpkg/web/cache"
	"github.com/itpkg/web/i18n"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

//Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init config file",
			Flags:   []cli.Flag{ENV},
			Action: EnvAction(func(env string, _ *cli.Context) error {
				fn := fmt.Sprintf("%s.toml", env)
				if _, err := os.Stat(fn); err == nil {
					return fmt.Errorf("file %s already exists", fn)
				}

				sec, err := web.Random(512)
				if err != nil {
					return err
				}
				return web.Store(fn, &Config{
					Secrets: web.ToBase64(sec),
					HTTP: HTTP{
						Host: "localhost",
						Port: 3000,
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
			Action: IocAction(func(mux *martini.ClassicMartini, _ *cli.Context) error {
				cfg := mux.Injector.Get(reflect.TypeOf((*Config)(nil))).Interface().(*Config)
				if !cfg.IsProduction() {
					mux.Use(cors.Allow(&cors.Options{
						AllowOrigins:     []string{"*"},
						AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
						AllowHeaders:     []string{"Origin", "Authorization"},
						ExposeHeaders:    []string{"Content-Length"},
						AllowCredentials: true,
					}))
				}

				mux.Use(i18n.LangHandler)
				mux.Use(strict.Strict)
				mux.Use(render.Renderer(render.Options{
					Layout:     "layout",
					Extensions: []string{".html"},
				}))

				web.Loop(func(en web.Engine) error {
					en.Mount(mux)
					return nil
				})
				mux.RunOnAddr(fmt.Sprintf(":%d", cfg.HTTP.Port))
				return nil
			}),
		},
		{
			Name:    "routers",
			Aliases: []string{"ro"},
			Usage:   "print out all defined routes in match order, with names",
			Flags:   []cli.Flag{ENV},
			Action: func(c *cli.Context) {
				mux := martini.Classic()
				web.Loop(func(en web.Engine) error {
					en.Mount(mux)
					return nil
				})
				for _, r := range mux.Router.All() {
					fmt.Printf("%s\t%s\n", r.Method(), r.Pattern())
				}
			},
		},
		{
			Name:    "cache",
			Aliases: []string{"c"},
			Usage:   "cache operations",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list all cache items",
					Flags:   []cli.Flag{ENV},
					Action: InvokeAction(func(cp cache.Provider) error {
						keys, err := cp.Status()
						if err != nil {
							return err
						}
						for k, v := range keys {
							fmt.Printf("%s\t%d\n", k, v)
						}
						return nil
					}),
				},
				{
					Name:    "delete",
					Aliases: []string{"d"},
					Usage:   "delete item from cache",
					Flags: []cli.Flag{
						ENV,
						cli.StringFlag{
							Name:  "key, k",
							Value: "",
							Usage: "cache item's key",
						},
					},
					Action: IocAction(func(mux *martini.ClassicMartini, ctx *cli.Context) error {
						k := ctx.String("key")
						if k == "" {
							return errors.New("key mustn't null")
						}
						_, err := mux.Invoke(func(cp cache.Provider) {
							cp.Del(k)
						})
						return err
					}),
				},
				{
					Name:    "clear",
					Aliases: []string{"c"},
					Usage:   "delete all items from cache",
					Flags:   []cli.Flag{ENV},
					Action: InvokeAction(func(cp cache.Provider) error {
						return cp.Clear()
					}),
				},
			},
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
