package postgres

import (
	"testing"
	"totraz_store/pkg/config"
)

func TestConnect(t *testing.T) {
	cfg := config.Postgres{
		Name:     "chat",
		Driver:   "postgres",
		Host:     "localhost",
		Port:     1234,
		Password: "test",
		User:     "totalim_chat",
		SslMode:  "disable",
	}

	_, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

}
