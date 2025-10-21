package version

import "time"

type Versions struct {
	ID        int         `db:"id" json:"id"`
	Version   string      `db:"version" json:"version"`
	Up        interface{} `db:"up" json:"up"`     // jsonb -> map or struct
	Down      interface{} `db:"down" json:"down"` // jsonb -> map or struct
	State     string      `db:"state" json:"state"`
	CreatedAt time.Time   `db:"created_at" json:"createdAt"`
	AppliedAt *time.Time  `db:"applied_at" json:"appliedAt,omitempty"`
	ProjectID int         `db:"project_id" json:"projectId"`
}

type VersionAudit struct {
	ID        int       `db:"id" json:"id"`
	VersionID int       `db:"version_id" json:"versionId"`
	AppliedAt time.Time `db:"applied_at" json:"appliedAt"`
	AppliedBy int       `db:"applied_by" json:"appliedBy"`
	Notes     string    `db:"notes" json:"notes"`
}
