package logging

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go_learn/017.utils/pathutils"
	"path"
	"time"
)

var Log *logrus.Logger

//NewLoggerWithRotate 返回日志对象, 日志路径为access.log, 且会按照日期切割, 单例模式
func NewLoggerWithRotate() (*logrus.Logger, error) {
	if Log != nil {
		return Log, nil
	}

	dir, err := pathutils.ProjectAbsPath()
	if err != nil {
		return nil, err
	}
	logPath := path.Join(dir, "access.log")
	writer, _ := rotatelogs.New(
		logPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(logPath),             // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(24*3600*time.Second),   // 文件最大保存时间
		rotatelogs.WithRotationTime(600*time.Second), // 日志切割时间间隔
	)

	levels := logrus.AllLevels
	pathMap := lfshook.WriterMap{}
	for _, level := range levels {
		pathMap[level] = writer
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Log, nil
}
