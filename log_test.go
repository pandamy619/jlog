package jlog

import (
	"testing"
)

func TestSimpleLogInfo(t *testing.T) {
	log := Init("Logging", "log", "tmp")
	log.Info(
		Fields{
			"func":    "TestSimpleLogInfo",
			"message": "info",
		},
	)
	log.Info("message ", "info")
	log.Info("message info")

}
func TestSimpleLogWarning(t *testing.T) {
	log := Init("Logging", "log", "tmp")
	log.Warning(
		Fields{
			"func":    "TestSimpleLogWarning",
			"message": "warning",
		},
	)
	log.Warning("message ", "warning")
	log.Warning("message warning")

}
