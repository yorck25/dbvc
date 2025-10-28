package projects

import (
	"backend/core"
	"strconv"
)

func HandleCreateProject(ctx *core.WebContext) error {
	userID := ctx.GetUserId()

	var cpr CreateProjectRequest
	if err := ctx.Bind(&cpr); err != nil {
		return ctx.BadRequest("invalid input")
	}

	repo := NewRepository(ctx)

	project, err := repo.CreateProject(cpr, userID)
	if err != nil {
		return ctx.InternalError("Error Create Project Metadata: " + err.Error())
	}

	err = repo.CreateProjectCredentials(cpr.Credentials, project.ID)
	if err != nil {
		return ctx.InternalError("Error Create Project Credentials: " + err.Error())
	}

	err = repo.CreateProjectMembers(cpr.Members, project.ID)
	if err != nil {
		return ctx.InternalError("Error Create Project Members: " + err.Error())
	}

	return ctx.Sucsess(project)
}

func HandleGetAllProjects(ctx *core.WebContext) error {
	_ = ctx.GetUserId()

	repo := NewRepository(ctx)
	projects, err := repo.GetAllProjects()
	if err != nil {
		return ctx.InternalError(err.Error())
	}
	return ctx.Sucsess(projects)
}

func HandleGetProjectByID(ctx *core.WebContext) error {
	_ = ctx.GetUserId()

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
	_ = ctx.GetUserId()

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
	_ = ctx.GetUserId()

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
	_ = ctx.GetUserId()

	repo := NewRepository(ctx)
	projects, err := repo.GetActiveProjects()
	if err != nil {
		return ctx.InternalError(err.Error())
	}
	return ctx.Sucsess(projects)
}
