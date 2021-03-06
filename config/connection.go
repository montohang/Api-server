package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	logs "github.com/MaulIbra/logs_module"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(env *Env) *sql.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.SchemaName)
	db, _ := sql.Open(env.Driver, dataSource)

	if err := db.Ping(); err != nil {
		panic(err)
		logs.ErrorLogger.Println(err)
		return nil
	}
	return db
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}