package projects

import (
	"backend/core"
	"strconv"
)

func HandleCreateProject(ctx *core.WebContext) error {
	var project Projects
	if err := ctx.Bind(&project); err != nil {
		return ctx.BadRequest("invalid input")
	}

	repo := NewRepository(ctx)
	if err := repo.CreateProject(&project); err != nil {
		return ctx.InternalError(err.Error())
	}

	return ctx.Sucsess(project)
}

func HandleGetAllProjects(ctx *core.WebContext) error {
	repo := NewRepository(ctx)
	projects, err := repo.GetAllProjects()
	if err != nil {
		return ctx.InternalError(err.Error())
	}
	return ctx.Sucsess(projects)
}

func HandleGetProjectByID(ctx *core.WebContext) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.BadRequest("invalid project id")
	}

	repo := NewRepository(ctx)
	project, err := repo.GetProjectByID(id)
	if err != nil {
		return ctx.NotFound("project not found")
	}

	return ctx.Sucsess(project)
}

func HandleUpdateProject(ctx *core.WebContext) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.BadRequest("invalid project id")
	}

	var project Projects
	if err := ctx.Bind(&project); err != nil {
		return ctx.BadRequest("invalid input")
	}
	project.ID = id

	repo := NewRepository(ctx)
	if err := repo.UpdateProject(&project); err != nil {
		return ctx.InternalError(err.Error())
	}

	return ctx.Sucsess(project)
}

func HandleDeleteProject(ctx *core.WebContext) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.BadRequest("invalid project id")
	}

	repo := NewRepository(ctx)
	if err := repo.DeleteProject(id); err != nil {
		return ctx.InternalError(err.Error())
	}

	return ctx.Sucsess(map[string]string{"message": "project deleted"})
}

func HandleGetActiveProjects(ctx *core.WebContext) error {
	repo := NewRepository(ctx)
	projects, err := repo.GetActiveProjects()
	if err != nil {
		return ctx.InternalError(err.Error())
	}
	return ctx.Sucsess(projects)
}
