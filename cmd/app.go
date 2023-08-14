package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	app = &cli.App{
		Name:  "Digital-Wallet",
		Usage: "providing functionalities for a user wallet",
		Action: func(*cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{apiCommand},
	}
	banner = `▌║█║▌│║▌│║▌║▌█║ Digital-Wallet ▌│║▌║▌│║║▌█║▌║█`
)

func Run() error {
	fmt.Println(banner)
	return app.Run(os.Args)
}
