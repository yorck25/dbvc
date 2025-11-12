package connectors

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

func (p *PostgresConnector) Connect(connectionString string) (*sql.DB, error) {
	fmt.Println("Connecting to PostgreSQL...")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func (p *PostgresConnector) BuildConnectionString(projectID int, metaDB *sql.DB) (string, error) {
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
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		auth.DatabaseAuth["Host"],
		auth.DatabaseAuth["Port"],
		auth.DatabaseAuth["Username"],
		auth.DatabaseAuth["Password"],
		auth.DatabaseAuth["Database"],
	)
	return connectionString, nil
}

func (p PostgresConnector) GetVersionQuery() string {
	return "SELECT version();"
}

func (p *PostgresConnector) ExecuteQuery(projectID int, query string) (*sql.Rows, error) {
	conStr, err := p.BuildConnectionString(projectID, p.MetaDataClient)
	if err != nil {
		return nil, err
	}

	db, err := p.Connect(conStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db.Query(query)
}

func (p *PostgresConnector) GetDatabaseStructure(projectID int) (*DatabaseStructureResponse, error) {
	conStr, err := p.BuildConnectionString(projectID, p.MetaDataClient)
	if err != nil {
		return nil, err
	}

	db, err := p.Connect(conStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
        SELECT
            table_schema,
            table_name,
            column_name,
            data_type
        FROM information_schema.columns
        ORDER BY table_schema, table_name, ordinal_position;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type TableMap struct {
		TableName string
		Columns   []ColumnStructureResponse
	}

	type SchemaMap struct {
		SchemaName string
		Tables     map[string]*TableMap
	}

	schemaMap := make(map[string]*SchemaMap)

	for rows.Next() {
		var schemaName, tableName, columnName, dataType string
		if err := rows.Scan(&schemaName, &tableName, &columnName, &dataType); err != nil {
			return nil, err
		}

		if _, ok := schemaMap[schemaName]; !ok {
			schemaMap[schemaName] = &SchemaMap{
				SchemaName: schemaName,
				Tables:     make(map[string]*TableMap),
			}
		}

		schema := schemaMap[schemaName]
		if _, ok := schema.Tables[tableName]; !ok {
			schema.Tables[tableName] = &TableMap{
				TableName: tableName,
				Columns:   []ColumnStructureResponse{},
			}
		}

		schema.Tables[tableName].Columns = append(schema.Tables[tableName].Columns, ColumnStructureResponse{
			ColumnName: columnName,
			DataType:   dataType,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	response := DatabaseStructureResponse{}
	for _, schema := range schemaMap {
		schemaResponse := SchemaStructureResponse{
			SchemaName: schema.SchemaName,
			Tables:     []TableStructureResponse{},
		}
		for _, table := range schema.Tables {
			schemaResponse.Tables = append(schemaResponse.Tables, TableStructureResponse{
				TableName: table.TableName,
				Columns:   table.Columns,
			})
		}
		response.Schemas = append(response.Schemas, schemaResponse)
	}

	return &response, nil
}
