package user

import (
	"backend/core"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(ctx *core.WebContext) *Repository {
	return &Repository{db: ctx.GetDb()}
}

func (r *Repository) CreateUser(firstName *string, email string) (*User, error) {
	query := `
        INSERT INTO users (first_name, email)
        VALUES (:first_name, :email)
        RETURNING id, first_name, email, created_at, updated_at, active`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &User{}

	params := map[string]any{
		"first_name": firstName,
		"email":      email,
	}

	err = stmt.Get(user, params)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) CreateUserLogin(userID int, username, passwordHash string) error {
	query := `
        INSERT INTO user_login (user_id, username, password_hash, failed_attempts)
        VALUES (:user_id, :username, :password_hash, 0)`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	params := map[string]any{
		"user_id":       userID,
		"username":      username,
		"password_hash": passwordHash,
	}

	_, err = stmt.Exec(params)

	return err
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	query := `
        SELECT id, first_name, email, created_at, updated_at, active
        FROM users
        WHERE email = :email`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &User{}
	err = stmt.Get(user, map[string]any{"email": email})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserByID(id int) (*User, error) {
	query := `
        SELECT id, first_name, email, created_at, updated_at, active
        FROM users
        WHERE id = :id`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &User{}
	err = stmt.Get(user, map[string]any{"id": id})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserLoginByUsername(username string) (*UserLogin, error) {
	query := `
        SELECT user_id, username, password_hash, last_login_at, failed_attempts
        FROM user_login
        WHERE username = :username`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	login := &UserLogin{}
	err = stmt.Get(login, map[string]any{"username": username})

	if err != nil {
		return nil, err
	}

	return login, nil
}

func (r *Repository) GetUserLoginByUserID(userID int) (*UserLogin, error) {
	query := `
        SELECT user_id, username, password_hash, last_login_at, failed_attempts
        FROM user_login
        WHERE user_id = :user_id`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	login := &UserLogin{}
	err = stmt.Get(login, map[string]any{"user_id": userID})

	if err != nil {
		return nil, err
	}

	return login, nil
}

func (r *Repository) UpdateLastLogin(userID int) error {
	query := `
        UPDATE user_login
        SET last_login_at = :last_login_at, failed_attempts = 0
        WHERE user_id = :user_id`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(map[string]any{
		"last_login_at": time.Now(),
		"user_id":       userID,
	})

	return err
}

func (r *Repository) IncrementFailedAttempts(username string) error {
	query := `
        UPDATE user_login
        SET failed_attempts = failed_attempts + 1
        WHERE username = :username`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(map[string]any{"username": username})
	return err
}

func (r *Repository) UpdatePassword(userID int, passwordHash string) error {
	query := `
        UPDATE user_login
        SET password_hash = :password_hash
        WHERE user_id = :user_id`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(map[string]any{
		"password_hash": passwordHash,
		"user_id":       userID,
	})

	return err
}

func (r *Repository) UpdateUser(userID int, firstName *string, email *string) (*User, error) {
	query := `
        UPDATE users 
        SET updated_at = :updated_at`

	params := map[string]any{
		"updated_at": time.Now(),
		"user_id":    userID,
	}

	if firstName != nil {
		query += `, first_name = :first_name`
		params["first_name"] = firstName
	}

	if email != nil {
		query += `, email = :email`
		params["email"] = email
	}

	query += ` WHERE id = :user_id RETURNING id, first_name, email, created_at, updated_at, active`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &User{}
	err = stmt.Get(user, params)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) CheckUsernameExists(username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM user_login WHERE username = :username)`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var exists bool
	err = stmt.Get(&exists, map[string]any{"username": username})

	return exists, err
}

func (r *Repository) CheckEmailExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = :email)`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var exists bool
	err = stmt.Get(&exists, map[string]any{"email": email})

	return exists, err
}
