package common

import "testing"

func TestGetLogger(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		logger := GetLogger()
		if logger == nil {
			t.Error("logger should not be nil")
		}
	})
}
