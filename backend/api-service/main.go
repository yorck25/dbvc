package main

import (
	"backend/core"
	"backend/projects"
	"backend/release"
	"backend/user"
	"backend/version"
	"backend/worker"

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

	app.POST("/auth/login", user.HandleLogin)
	app.POST("/auth/register", user.HandleRegister)
	app.GET("/auth/me", user.HandleGetProfile)

	app.GET("/users/search", user.HandleSearchMembers)

	app.GET("/projects", projects.GetAllProjectsWithUsersForUser)
	app.GET("/projects/all", projects.HandleGetAllProjects)
	app.GET("/projects/active", projects.HandleGetActiveProjects)
	app.GET("/projects/:id", projects.HandleGetProjectByID)
	app.POST("/projects", projects.HandleCreateProject)
	app.PUT("/projects/:id", projects.HandleUpdateProject)
	app.DELETE("/projects/:id", projects.HandleDeleteProject)

	app.GET("/release/project/all", release.HandleGetReleasesForProject)
	app.GET("/release/project/latest", release.HandleGetLatestReleasesForProject)

	app.POST("/version/table/create", version.HandleCreateTable)

	app.GET("/upgrade", worker.HandleUpgrade)
	app.GET("/connection-types", worker.HandleGetConnectionTypes)

	app.Logger.Fatal(app.Start("0.0.0.0:8080"))
}

func defaultUrl(ctx *core.WebContext) error {
	return ctx.Sucsess("This is the backend")
}

func healthUrl(ctx *core.WebContext) error {
	return ctx.Sucsess("OK")
}
