package config_test

import (
	"testing"

	"github.com/webbergao1/go-toolkit/config"
)

func Test_LoadConfig(t *testing.T) {
	testfile := "./app_test.toml"
	conf, err := config.LoadConfig(testfile)
	if err != nil {
		t.Error(err)
	}
	t.Logf("config is: %#v", conf)
	v, ok := conf.(map[string]string)
	if !ok {
		t.Error(ok)
	}
	t.Logf("config is: %#v", v)
}
