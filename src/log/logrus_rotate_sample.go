package log

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

func initLog(log *logrus.Logger) {
	log.SetLevel(logrus.InfoLevel)
	log.SetNoLock()
	configLocalFilesystemLogger(log, "", "logrus_rotate_sample_test.log",
		time.Second * 10, time.Hour)
}

func configLocalFilesystemLogger(log *logrus.Logger, logPath string, logFileName string,
	maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath + ".%Y%m%d%H%M%S",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(2),
	)

	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	//switch level := GlobalConfig.LogConf.LogLevel; level {
	///*
	//如果日志级别不是debug就不要打印日志到控制台了
	// */
	//case "debug":
	//	log.SetLevel(log.DebugLevel)
	//	log.SetOutput(os.Stderr)
	//case "info":
	//	setNull()
	//	log.SetLevel(log.InfoLevel)
	//case "warn":
	//	setNull()
	//	log.SetLevel(log.WarnLevel)
	//case "error":
	//	setNull()
	//	log.SetLevel(log.ErrorLevel)
	//default:
	//	setNull()
	//	log.SetLevel(log.InfoLevel)
	//}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel: writer,
		logrus.WarnLevel: writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{
		DisableColors:true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp: true})

	log.AddHook(lfHook)
}

func SampleMainLogrusRotate() {

	fmt.Println("\n[logrus_rotate_sample]")

	log := logrus.New()
	initLog(log)

	// 分别构造两个不同的log，多用于在不同的package中设置不同的log
	log1 := log.WithFields(logrus.Fields{
		"k1": "v1",
	})
	log2 := log.WithFields(logrus.Fields{
		"k2": "v2",
	})
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			log1.Info("test 00" + fmt.Sprintf("%d", i))
		} else {
			log2.Info("test 00" + fmt.Sprintf("%d", i))
		}
		time.Sleep(time.Millisecond * 100)
	}
}