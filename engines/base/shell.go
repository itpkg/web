package base

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/itpkg/web"
)

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
	}
}
