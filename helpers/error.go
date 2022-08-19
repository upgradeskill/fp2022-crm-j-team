package helpers

import (
	"github.com/sirupsen/logrus"
)

func LogIfError(err error, message string) {
	if err != nil {
		defer logrus.Info(message)
		logrus.Fatal(err)
	}
}
