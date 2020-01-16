package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Out = os.Stdout
	Log.SetLevel(logrus.TraceLevel)
}

func Main() {
	Log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	otherfunc()
}
