package oauth

import (
	"github.com/codegangsta/cli"
	"github.com/go-martini/martini"
	"github.com/itpkg/web/engines/base"
	"github.com/itpkg/web/settings"
)

//Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
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
						base.ENV,
						cli.StringFlag{
							Name:  "file, f",
							Value: "google.json",
							Usage: "filename",
						},
					},
					Action: base.IocAction(func(mux *martini.ClassicMartini, ctx *cli.Context) error {
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
