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
	return db.Query(query)
}

func (m MSSQLConnector) GetVersionQuery() string {
	return "SELECT @@VERSION;"
}

func (m MSSQLConnector) GetDatabaseStructure(db *sql.DB) (interface{}, error) {
	return nil, fmt.Errorf("Not implemented")
}
