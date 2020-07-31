package main

import (
	"fmt"
	"github.com/fideism/golang-tools/logger"
	"github.com/sirupsen/logrus"
)

func main(){
	logger.NewEntry().WithFields(logrus.Fields{
		"key": "value",
	}).Info("test message")


	logger.SetOption(logger.Channel("test"))
	logger.Entry().Info("channel sms")

	logger.Entry().WithError(fmt.Errorf("teseddddddddddddddt")).Info("aaaa")

	logger.Entry().WithFields(logrus.Fields{
		"a" : 1,
	}).Info("test fields")
}
