package connectors

import (
	"database/sql"
	"fmt"
)

func (m MSSQLConnector) Connect() (*sql.DB, error) {
	fmt.Println("Connecting to MSSQL...")
	db, err := sql.Open("sqlserver", m.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

func (m MSSQLConnector) ExecuteQuery(db *sql.DB, query string) (*sql.Rows, error) {
	fmt.Println("Executing query on MSSQL...")
	return db.Query(query)
}
