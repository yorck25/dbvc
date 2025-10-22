package connectors

import (
	"database/sql"
	"fmt"
)

func (m MySQLConnector) Connect() (*sql.DB, error) {
	fmt.Println("Connecting to MySQL...")
	db, err := sql.Open("mysql", m.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

func (m MySQLConnector) ExecuteQuery(db *sql.DB, query string) (*sql.Rows, error) {
	return db.Query(query)
}

func (m MySQLConnector) GetVersionQuery() string {
	return "SELECT VERSION();"
}
