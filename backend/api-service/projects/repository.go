package projects

import (
	"backend/core"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(ctx *core.WebContext) *Repository {
	return &Repository{db: ctx.GetDb()}
}

func (r *Repository) GetAllProjects() ([]Projects, error) {
	var projects []Projects
	stmt, err := r.db.PrepareNamed(`SELECT * FROM projects ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&projects, map[string]any{})
	if err != nil {
		return nil, err
	}

	return projects, err
}

func (r *Repository) GetProjectByID(id int) ([]Projects, error) {
	var projects []Projects
	stmt, err := r.db.PrepareNamed(`SELECT * FROM projects WHERE id = :id ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}

	params := map[string]any{
		"id": id,
	}

	err = stmt.Select(&projects, params)
	if err != nil {
		return nil, err
	}

	return projects, err
}

func (r *Repository) GetActiveProjects() ([]Projects, error) {
	var projects []Projects
	stmt, err := r.db.PrepareNamed(`SELECT * FROM projects WHERE active = :active ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}

	params := map[string]any{
		"active": true,
	}

	err = stmt.Select(&projects, params)
	if err != nil {
		return nil, err
	}

	return projects, err
}

func (r *Repository) CreateProject(p *Projects) error {
	stmt, err := r.db.PrepareNamed(`
		INSERT INTO projects (owner_id, name, description, created_at, updated_at, active, visibility, connection_type)
		VALUES (:owner_id, :name, :description, NOW(), NOW(), :active, :visibility, :connection_type)
		RETURNING id`)
	if err != nil {
		return err
	}

	params := map[string]any{
		"owner_id":        p.OwnerID,
		"name":            p.Name,
		"description":     p.Description,
		"active":          p.Active,
		"visibility":      p.Visibility,
		"connection_type": p.ConnectionType,
	}

	return stmt.Get(&p.ID, params)
}

func (r *Repository) UpdateProject(p *Projects) error {
	stmt, err := r.db.PrepareNamed(`
		UPDATE projects
		SET name = :name,
			description = :description,
			updated_at = NOW(),
			active = :active,
			visibility = :visibility,
			connection_type = :connection_type
		WHERE id = :id`)
	if err != nil {
		return err
	}

	params := map[string]any{
		"id":              p.ID,
		"name":            p.Name,
		"description":     p.Description,
		"active":          p.Active,
		"visibility":      p.Visibility,
		"connection_type": p.ConnectionType,
	}

	_, err = stmt.Exec(params)
	return err
}

func (r *Repository) DeleteProject(id int) error {
	stmt, err := r.db.PrepareNamed(`DELETE FROM projects WHERE id = :id`)
	if err != nil {
		return err
	}

	params := map[string]any{
		"id": id,
	}

	_, err = stmt.Exec(params)
	return err
}
