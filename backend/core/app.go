package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type App struct {
	*echo.Echo
	Ctx *AppContext
}

func InitApp() (*App, error) {
	ctx := &AppContext{}

	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	ctx.config = config

	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/bike")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	ctx.db = db

	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.Use(CreateCtx(ctx))

	return &App{Echo: e, Ctx: ctx}, nil
}

type HandlerFunc func(*WebContext) error

func (f HandlerFunc) Handle(ctx *WebContext) error {
	return f(ctx)
}

func wrapHandler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*WebContext)
		return h.Handle(ctx)
	}
}

func (a *App) Group(prefix string, m ...echo.MiddlewareFunc) *Group {
	g := a.Echo.Group(prefix, m...)
	return &Group{Group: g}
}

type Group struct {
	*echo.Group
}

func (g *Group) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(http.MethodGet, path, wrapHandler(h), m...)
}

func (g *Group) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(http.MethodPost, path, wrapHandler(h), m...)
}

func (g *Group) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(http.MethodPut, path, wrapHandler(h), m...)
}

func (g *Group) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(http.MethodDelete, path, wrapHandler(h), m...)
}
