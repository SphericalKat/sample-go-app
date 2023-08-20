package config

import (
	"log/slog"
	"sample/ent"

	_ "github.com/mattn/go-sqlite3"
)

func ProvideSqliteDB(conf *Config) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", conf.DatabaseURL)
	if err != nil {
		slog.Error("error connecting to sqlite database", "err", err)
		return nil, err
	}

	return client, nil
}
