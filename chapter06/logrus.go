package main

import "os"
import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout
	log.WithFields(logrus.Fields{}).Trace("TraceLog")
	log.WithFields(logrus.Fields{}).Info("InfoLog")
	log.WithFields(logrus.Fields{}).Error("ErrorLog")

}
