package jlog

import (
	"testing"
)

func TestSimpleLog(t *testing.T) {
	log := Init("Logging", "log", "tmp")
	log.Info("Info message")
	log.Warning("Warning message")
}
