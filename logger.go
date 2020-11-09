package main

import (
	"path"
	"time"
	"github.com/sirupsen/logrus"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var accessLog = logrus.New()
var errorLog = logrus.New()

func init() {
	dir := config.LogPath
	suffix := ".%Y%m%d%H%M"
	// 日志轮转
	filename := path.Join(dir, "access.log")
	accessWriter, _ := rotatelogs.New(
		filename + suffix,
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(7 * 24 * time.Hour),
		rotatelogs.WithRotationTime(24 * time.Hour),
	)
	accessLog.SetOutput(accessWriter)
	accessLog.SetFormatter(&logrus.JSONFormatter{})

	filename = path.Join(dir, "error.log")
	errorWriter, _ := rotatelogs.New(
		filename + suffix,
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(7 * 24 * time.Hour),
		rotatelogs.WithRotationTime(24 * time.Hour),
	)
	errorLog.SetOutput(errorWriter)
	errorLog.SetFormatter(&logrus.JSONFormatter{})
}
