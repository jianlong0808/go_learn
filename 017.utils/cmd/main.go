package main

import (
	"go_learn/017.utils/logging"
)

//Logging 测试日志生成
func Logging() {
	log, err := logging.NewLoggerWithRotate()
	if err != nil {
		panic("err")
	}
	log.Info("test")
}

func main() {
	Logging()
}
