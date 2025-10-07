package main

import (
	"backend/core"
	"github.com/labstack/echo/v4/middleware"
)

type UserClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app, err := core.InitApp()

	if err != nil {
		panic(err)
	}

	app.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	app.GET("/", defaultUrl)
	app.GET("/health", healthUrl)

	app.Logger.Fatal(app.Start("0.0.0.0:8080"))
}

func defaultUrl(ctx *core.WebContext) error {
	return ctx.Sucsess("This is the backend")
}

func healthUrl(ctx *core.WebContext) error {
	return ctx.Sucsess("OK")
}
