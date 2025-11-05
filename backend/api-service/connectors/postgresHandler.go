package connectors

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func (p PostgresConnector) Connect() (*sql.DB, error) {
	fmt.Println("Connecting to PostgreSQL...")
	db, err := sql.Open("postgres", p.ConnectionString)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

func (p PostgresConnector) ExecuteQuery(db *sql.DB, query string) (*sql.Rows, error) {
	return db.Query(query)
}

func (p PostgresConnector) GetVersionQuery() string {
	return "SELECT version();"
}

func (m PostgresConnector) GetDatabaseStructure(db *sql.DB) (interface{}, error) {
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

	type Column struct {
		ColumnName string `json:"columnName"`
		DataType   string `json:"dataType"`
	}

	type Table struct {
		TableName string   `json:"tableName"`
		Columns   []Column `json:"columns"`
	}

	type Schema struct {
		SchemaName string            `json:"schemaName"`
		Tables     map[string]*Table `json:"tables"`
	}

	structure := make(map[string]*Schema)

	for rows.Next() {
		var schemaName, tableName, columnName, dataType string
		if err := rows.Scan(&schemaName, &tableName, &columnName, &dataType); err != nil {
			return nil, err
		}

		if _, ok := structure[schemaName]; !ok {
			structure[schemaName] = &Schema{
				SchemaName: schemaName,
				Tables:     make(map[string]*Table),
			}
		}

		schema := structure[schemaName]
		if _, ok := schema.Tables[tableName]; !ok {
			schema.Tables[tableName] = &Table{
				TableName: tableName,
				Columns:   []Column{},
			}
		}

		schema.Tables[tableName].Columns = append(schema.Tables[tableName].Columns, Column{
			ColumnName: columnName,
			DataType:   dataType,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return structure, nil
}
