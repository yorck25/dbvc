package projects

import "time"

type Projects struct {
	ID             int       `db:"id" json:"id"`
	OwnerID        int       `db:"owner_id" json:"ownerId"`
	Name           string    `db:"name" json:"name"`
	Description    string    `db:"description" json:"description"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`
	Active         bool      `db:"active" json:"active"`
	Visibility     string    `db:"visibility" json:"visibility"`
	ConnectionType int       `db:"connection_type" json:"connectionType"`
}

type CreateProjectRequest struct {
	Name           string `db:"name" json:"name"`
	Description    string `db:"description" json:"description"`
	Visibility     string `db:"visibility" json:"visibility"`
	ConnectionType int    `db:"connection_type" json:"connectionType"`
}
