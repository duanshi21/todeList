package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var LoggerObj *logrus.Logger

func InitLog() {
	if LoggerObj != nil {
		src, _ := setOutputFile()
		LoggerObj.Out = src

		return
	}

	// 日志对象的实例化
	logger := logrus.New()
	src, _ := setOutputFile()
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05",
	})

	LoggerObj = logger
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err != nil {
		logFilePath = dir + "/logs/" // 日志输出的文件夹路径
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log" // 以日期来进行区分文件，放在对应的文件夹内
	fileName := path.Join(logFilePath, logFileName)  // 日志输出文件
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			return nil, err
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	return src, nil
}
