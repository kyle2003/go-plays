package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"pandora/services"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	services.Start()
}
