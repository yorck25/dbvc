package user

import "time"

type Users struct {
	ID        int       `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"firstName"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Active    bool      `db:"active" json:"active"`
}

type UserRole struct {
	ID        int    `db:"id" json:"id"`
	UserID    int    `db:"user_id" json:"userId"`
	ProjectID int    `db:"project_id" json:"projectId"`
	Role      string `db:"role" json:"role"`
}

type UserLogin struct {
	UserID         int        `db:"user_id" json:"userId"`
	Username       string     `db:"username" json:"username"`
	PasswordHash   string     `db:"password_hash" json:"passwordHash"`
	LastLoginAt    *time.Time `db:"last_login_at" json:"lastLoginAt,omitempty"`
	FailedAttempts int        `db:"failed_attempts" json:"failedAttempts"`
}
