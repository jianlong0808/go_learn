package logging

import (
	"testing"
)

func TestLogging(t *testing.T) {
	log, err := NewLoggerWithRotate()
	if err != nil {
		panic("err")
	}
	log.Info("test")
}
