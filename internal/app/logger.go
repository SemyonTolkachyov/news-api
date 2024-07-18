package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

// initLogger init default logger
func initLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{})
	if os.Getenv("MODE") == "DEBUG" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
