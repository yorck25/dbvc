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

func (m MySQLConnector) Disconnect() error {
	err := m.Client.Close()
	return err
}

func (m MySQLConnector) ExecuteQuery(query string) (*sql.Rows, error) {
	return m.Client.Query(query)
}

func (m MySQLConnector) GetVersionQuery() string {
	return "SELECT VERSION();"
}

func (m MySQLConnector) GetDatabaseStructure() (*DatabaseStructureResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}
