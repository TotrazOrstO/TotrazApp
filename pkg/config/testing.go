package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// КОСТЫЛЬ. получает путь до корня проекта
	basepath = filepath.Dir(strings.Replace(b, "/TotrazApp/pkg/config", "/TotrazApp", 1))
)

func TestConfigs(t *testing.T) Config {
	t.Helper()
	err := os.Chdir(basepath)
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	return cfg

}
