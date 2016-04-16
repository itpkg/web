package dict

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/itpkg/web/engines/base"
)

//Shell shell commands
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "dicts",
			Aliases: []string{"ds"},
			Usage:   "list dicts",
			Flags:   []cli.Flag{base.ENV},
			Action: base.InvokeAction(func(dp Provider) error {
				ds, er := dp.List()
				if er != nil {
					return er
				}
				fmt.Println("Dictionary's name\tWord count")
				for k, v := range ds {
					fmt.Printf("%s\t%d\n", k, v)
				}
				return nil
			}),
		},
	}
}
