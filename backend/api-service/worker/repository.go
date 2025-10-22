package worker

import (
	"backend/core"
	"backend/projects"
	"backend/version"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(ctx *core.WebContext) *Repository {
	return &Repository{db: ctx.GetDb()}
}

func (r *Repository) GetProjectDetails(id int) (projects.Projects, error) {
	var project projects.Projects
	stmt, err := r.db.PrepareNamed("SELECT * FROM projects WHERE id = :id")
	if err != nil {
		return project, err
	}

	params := map[string]any{
		"id": id,
	}

	err = stmt.Get(&project, params)
	if err != nil {
		return project, err
	}

	return project, nil
}

func (r *Repository) GetConnectionTypeObject(id int) (ConnectionType, error) {
	var connectionType ConnectionType

	stmt, err := r.db.PrepareNamed("SELECT * FROM connection_types WHERE id = :id")
	if err != nil {
		return connectionType, err
	}

	params := map[string]any{
		"id": id,
	}

	err = stmt.Get(&connectionType, params)
	if err != nil {
		return connectionType, err
	}

	return connectionType, err
}

func (r *Repository) GetVersionsSinceLatestRelease(projectId int, latestVersionId int) ([]version.Versions, error) {
	var versions []version.Versions

	stmt, err := r.db.PrepareNamed("SELECT * FROM versions WHERE project_id = :projectId AND id > :versionId ORDER BY id")
	if err != nil {
		return versions, err
	}

	params := map[string]any{
		"projectId": projectId,
		"versionId": latestVersionId,
	}

	err = stmt.Select(&versions, params)
	if err != nil {
		return versions, err
	}

	return versions, err
}
