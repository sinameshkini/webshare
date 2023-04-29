package main

import (
	"github.com/sinameshkini/webshare/cmd"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}
