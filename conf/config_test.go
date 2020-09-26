package conf

import "testing"

func TestConfig(t *testing.T) {

	level := map[string]int64{
		"INFO":  10,
		"WARN":  100,
		"ERROR": 1000,
	}
	outputPaths := "/tmp/log"
	config := NewSausageLogConf(level, outputPaths)
	t.Log(config.GetLevel("INFO"))
	t.Log(config.GetOutputPaths())
	// config.Level =
}
