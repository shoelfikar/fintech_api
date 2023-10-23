package main

import (
	"net/http"

	"github.com/shoelfikar/kreditplus/app"
	"github.com/shoelfikar/kreditplus/db/drivers"
	"github.com/shoelfikar/kreditplus/pkg"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	pkg.NewViperLoad()

	DB := drivers.NewMYSQL()

	router := app.NewRouter(DB)

	server := http.Server{
		Addr: ":" +pkg.GetViperEnvVariable("PORT"),
		Handler: router,
	}

	server.ListenAndServe()
}