package web

import (
	"os"

	"github.com/codegangsta/cli"
)

//Main entry
func Main() error {
	app := cli.NewApp()
	app.Name = "itpkg"
	app.Usage = "IT-PACKAGE web framework"
	app.Version = "v20160401"
	app.Commands = make([]cli.Command, 0)
	for _, en := range engines {
		cmd := en.Shell()
		app.Commands = append(app.Commands, cmd...)
	}

	return app.Run(os.Args)
}
