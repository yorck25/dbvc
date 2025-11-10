package connectors

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func (p *PostgresConnector) Connect() (*sql.DB, error) {
	fmt.Println("Connecting to PostgreSQL...")
	db, err := sql.Open("postgres", p.ConnectionString)
	if err != nil {
		return nil, err
	}

	p.Client = db

	return db, db.Ping()
}

func (p PostgresConnector) Disconnect() error {
	err := p.Client.Close()
	return err
}

func (p PostgresConnector) ExecuteQuery(query string) (*sql.Rows, error) {
	return p.Client.Query(query)
}

func (p PostgresConnector) GetVersionQuery() string {
	return "SELECT version();"
}

func (m PostgresConnector) GetDatabaseStructure() (*DatabaseStructureResponse, error) {
	query := `
        SELECT
            TABLE_SCHEMA,
            TABLE_NAME,
            COLUMN_NAME,
            DATA_TYPE
        FROM INFORMATION_SCHEMA.COLUMNS
        ORDER BY TABLE_SCHEMA, TABLE_NAME, ORDINAL_POSITION;
    `

	rows, err := m.Client.Query(query)
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

		table := schema.Tables[tableName]

		table.Columns = append(table.Columns, ColumnStructureResponse{
			ColumnName: columnName,
			DataType:   dataType,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	response := DatabaseStructureResponse{
		Schemas: []SchemaStructureResponse{},
	}

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
