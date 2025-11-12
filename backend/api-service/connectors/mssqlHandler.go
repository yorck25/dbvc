package connectors

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"

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

func (m *MSSQLConnector) BuildConnectionString(projectID int, metaDB *sqlx.DB) (string, error) {
	query := `
		SELECT database_auth
		FROM projects_credentials
		WHERE project_id = $1
	`
	var authJSON string
	if err := metaDB.QueryRow(query, projectID).Scan(&authJSON); err != nil {
		return "", err
	}

	var auth DatabaseAuth
	if err := json.Unmarshal([]byte(authJSON), &auth.DatabaseAuth); err != nil {
		return "", fmt.Errorf("failed to parse database_auth JSON: %v", err)
	}

	decode := func(s string) string {
		if s == "" {
			return ""
		}
		data, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return s
		}
		return string(data)
	}

	connectionString := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		decode(auth.DatabaseAuth["Username"]),
		decode(auth.DatabaseAuth["Password"]),
		decode(auth.DatabaseAuth["Host"]),
		decode(auth.DatabaseAuth["Port"]),
		decode(auth.DatabaseAuth["databaseName"]),
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
