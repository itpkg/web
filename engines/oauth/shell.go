package oauth

import (
	"errors"
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/go-martini/martini"
	"github.com/itpkg/web/config"
	"github.com/itpkg/web/settings"
	"github.com/itpkg/web/token"
	"github.com/jinzhu/gorm"
)

//Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "users",
			Aliases: []string{"us"},
			Usage:   "users manage",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list all users",
					Flags: []cli.Flag{
						config.ENV,
					},
					Action: config.DbAction(func(db *gorm.DB, ctx *cli.Context) error {
						var us []User
						if err := db.Select([]string{"uid", "name", "email"}).Order("last_sign_in DESC").Find(&us).Error; err != nil {
							return err
						}
						fmt.Println("UID\t\t\t\t\tUSER")
						for _, u := range us {
							fmt.Printf("%s\t%s<%s>\n", u.UID, u.Name, u.Email)
						}
						return nil
					}),
				},
				{
					Name:    "role",
					Aliases: []string{"r"},
					Usage:   "role manage",
					Flags: []cli.Flag{
						config.ENV,
						cli.StringFlag{
							Name:  "name, n",
							Value: "",
							Usage: "role name(not empty)",
						},
						cli.StringFlag{
							Name:  "uid, u",
							Value: "",
							Usage: "user's uid(not empty)",
						},
						cli.IntFlag{
							Name:  "year, y",
							Value: 10,
							Usage: "years",
						},
						cli.BoolFlag{
							Name:  "deny, d",
							Usage: "remove role from user",
						},
					},
					Action: config.DbAction(func(db *gorm.DB, ctx *cli.Context) error {
						role := ctx.String("name")
						user := ctx.String("uid")
						deny := ctx.Bool("deny")
						year := ctx.Int("year")

						if role == "" {
							return errors.New("role's name mustn't empty")
						}
						if user == "" {
							return errors.New("uid mustn't empty")
						}
						dao := Dao{Db: db}
						r, err := dao.Role(role, "-", 0)
						if err != nil {
							return err
						}
						u, err := dao.GetUser(user)
						if err != nil {
							return err
						}
						if deny {
							err = dao.Deny(r.ID, u.ID)
						} else {
							err = dao.Allow(r.ID, u.ID, time.Hour*24*365*time.Duration(year))
						}

						return err
					}),
				},
			},
		},
		{
			Name:    "tokens",
			Aliases: []string{"tk"},
			Usage:   "online tokens ",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "list all tokens",
					Flags: []cli.Flag{
						config.ENV,
					},
					Action: config.InvokeAction(func(jwt *token.Jwt) error {
						ks, er := jwt.Provider.All()
						if er != nil {
							return er
						}
						fmt.Println("KID\t\t\t\tTTL")
						for k, v := range ks {
							fmt.Printf("%s\t%d\n", k, v)
						}
						return nil
					}),
				},
				{
					Name:    "clear",
					Aliases: []string{"c"},
					Usage:   "clear all tokens",
					Flags: []cli.Flag{
						config.ENV,
					},
					Action: config.InvokeAction(func(jwt *token.Jwt) error {
						return jwt.Provider.Clear()
					}),
				},
			},
		},
		{
			Name:    "oauth",
			Aliases: []string{"o"},
			Usage:   "oauth settings",
			Subcommands: []cli.Command{
				{
					Name:    "google",
					Aliases: []string{"g"},
					Usage:   "import google+ credentials",
					Flags: []cli.Flag{
						config.ENV,
						cli.StringFlag{
							Name:  "file, f",
							Value: "google.json",
							Usage: "filename",
						},
					},
					Action: config.IocAction(func(mux *martini.ClassicMartini, ctx *cli.Context) error {
						gfg, err := ReadGoogle(ctx.String("file"))
						if err != nil {
							return err
						}
						mux.Invoke(func(dao settings.Provider) {
							dao.Set("oauth.google", gfg, true)
						})
						return nil
					}),
				},
			},
		},
	}
}
