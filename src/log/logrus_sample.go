package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func getLog1(log *logrus.Logger) *logrus.Logger {
	log.WithFields(logrus.Fields{
		"animal": "wangqiang",
	})
	return log
}

func getLog2(log *logrus.Logger) *logrus.Logger {
	log.WithFields(logrus.Fields{
		"animal": "male",
	})
	return log
}

func writeLog1(log *logrus.Logger, content string) {
	log.WithFields(logrus.Fields{
		"name": "wq",
	}).Info(content)
}

func writeLog2(log *logrus.Logger, content string) {
	log.WithFields(logrus.Fields{
		"name": "123",
	}).Info(content)
}

func SampleMainLogrus()  {

	fmt.Println("\n[logrus_sample]")

	var log = logrus.New()
	// 通过标准输出写日志
	log.Out = os.Stdout

	getLog1(log).Info("A group of walrus emerges from the ocean")
	getLog2(log).Info("A group of walrus emerges from the ocean")

	writeLog1(log, "hello, world")
	writeLog2(log, "hello, world")
}