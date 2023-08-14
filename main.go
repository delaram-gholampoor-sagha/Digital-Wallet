package main

import (
	"os"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/cmd"
	"github.com/labstack/gommon/log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
