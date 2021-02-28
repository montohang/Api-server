package config

import (
	"api_server/utils"
	"os"
)


type Env struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	SchemaName string
	Driver     string
}

func NewEnv() *Env {
	env := Env{}
	env.DBUser = os.Getenv("DB_USER")
	if os.Getenv("DB_USER") == "" {
		env.DBUser = utils.GetEnv("DB_USER", "root")
	}

	env.DBPassword = os.Getenv("DB_PASSWORD")
	if os.Getenv("DB_PASSWORD") == "" {
		env.DBPassword = utils.GetEnv("DB_PASSWORD", "ini-password-2020")
	}

	env.DBHost = os.Getenv("DB_HOST")
	if os.Getenv("DB_HOST") == "" {
		env.DBHost = utils.GetEnv("DB_HOST", "127.0.0.1")
	}

	env.DBPort = os.Getenv("DB_PORT")
	if os.Getenv("DB_PORT") == "" {
		env.DBPort = utils.GetEnv("DB_PORT", "3306")
	}

	env.SchemaName = os.Getenv("DB_SCHEMA")
	if os.Getenv("DB_SCHEMA") == "" {
		env.SchemaName = utils.GetEnv("DB_SCHEMA", "db_user")
	}

	env.Driver = os.Getenv("DB_DRIVER")
	if os.Getenv("DB_DRIVER") == "" {
		env.Driver = utils.GetEnv("DB_DRIVER", "mysql")
	}

	return &env
}