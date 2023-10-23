package drivers

import (
	"database/sql"
	"time"

	"github.com/shoelfikar/kreditplus/pkg"
)

func NewMYSQL() *sql.DB {
	db, err := sql.Open("mysql", pkg.GetViperEnvVariable("DB_URL"))
	pkg.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	pkg.DefaultLoggingDebug("success connect to mysql database")

	return db
}