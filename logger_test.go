package log

import "testing"

func  TestNewLogger(t *testing.T) {
	logger := NewLogger("", "debug")

	logger.Info(" Hello World Logger !")
}