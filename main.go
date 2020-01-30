package main

import (
	l "go-playground/logging"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func main() {
	l.Log.Info("Tester")
	l.InitLogging()
}
