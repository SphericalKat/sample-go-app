package config

import (
	"log/slog"
	"os"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	DatabaseURL    string `koanf:"DATABASE_URL"`
	DatabaseDriver string `koanf:"DATABASE_DRIVER"`
}

var Conf *Config

var k = koanf.New(".")

func ProvideConfig() *Config {
	if Conf == nil {
		k.Load(confmap.Provider(map[string]any{
			"DATABASE_URL":    "file:ent?mode=memory&cache=shared&_fk=1",
			"DATABASE_DRIVER": "sqlite3",
		}, ""), nil)

		if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
			slog.Info("unable to find env file", "err", err)
			slog.Info("falling back to env variables")
		}

		if err := k.Load(env.Provider("", ".", nil), nil); err != nil {
			slog.Error("error loading config", "err", err)
			os.Exit(1)
		}

		Conf = &Config{}
		err := k.Unmarshal("", Conf)
		if err != nil {
			slog.Error("error loading config", "err", err)
			os.Exit(1)
		}
	}

	return Conf
}
