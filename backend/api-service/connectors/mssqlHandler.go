package connectors

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func (m *MSSQLConnector) Connect(connectionString string) (*sql.DB, error) {
	fmt.Println("Connecting to MSSQL...")
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func (m *MSSQLConnector) BuildConnectionString(projectID int, metaDB *sql.DB) (string, error) {
	query := `
		SELECT database_auth
		FROM project_credentials
		WHERE project_id = $1
	`
	var authJSON string
	if err := metaDB.QueryRow(query, projectID).Scan(&authJSON); err != nil {
		return "", err
	}

	var auth DatabaseAuth
	if err := json.Unmarshal([]byte(authJSON), &auth); err != nil {
		return "", fmt.Errorf("failed to parse database_auth JSON: %v", err)
	}

	connectionString := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		auth.DatabaseAuth["Host"],
		auth.DatabaseAuth["Port"],
		auth.DatabaseAuth["Username"],
		auth.DatabaseAuth["Password"],
		auth.DatabaseAuth["Database"],
	)
	return connectionString, nil
}

func (m *MSSQLConnector) ExecuteQuery(projectId int, query string) (*sql.Rows, error) {
	conStr, err := m.BuildConnectionString(projectId, m.MetaDataClient)
	if err != nil {
		return nil, err
	}

	db, err := m.Connect(conStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db.Query(query)
}

func (m MSSQLConnector) GetVersionQuery() string {
	return "SELECT @@VERSION;"
}

func (m MSSQLConnector) GetDatabaseStructure(projectId int) (*DatabaseStructureResponse, error) {
	conStr, err := m.BuildConnectionString(projectId, m.MetaDataClient)
	if err != nil {
		return nil, err
	}

	db, err := m.Connect(conStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
        SELECT
            TABLE_SCHEMA,
            TABLE_NAME,
            COLUMN_NAME,
            DATA_TYPE
        FROM INFORMATION_SCHEMA.COLUMNS
        ORDER BY TABLE_SCHEMA, TABLE_NAME, ORDINAL_POSITION;
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseDatabaseStructure(rows)
}
