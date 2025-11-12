package connectors

import (
	"database/sql"
)

type DBConnector interface {
	Connect(connectionString string) (*sql.DB, error)
	Disconnect() error
	ExecuteQuery(projectID int, query string) (*sql.Rows, error)
	GetVersionQuery() string
	GetDatabaseStructure(projectID int) (*DatabaseStructureResponse, error)
	BuildConnectionString(projectID int, metaDB *sql.DB) (string, error)
}

type DatabaseAuth struct {
	DatabaseAuth map[string]string `json:"databaseAuth"`
}

type PostgresConnector struct {
	MetaDataClient   *sql.DB
	ConnectionString string
}

type MySQLConnector struct {
	MetaDataClient   *sql.DB
	ConnectionString string
}

type MSSQLConnector struct {
	MetaDataClient   *sql.DB
	ConnectionString string
}

type DatabaseStructureResponse struct {
	Schemas []SchemaStructureResponse `json:"schemas"`
}

type SchemaStructureResponse struct {
	SchemaName string                   `json:"schemaName"`
	Tables     []TableStructureResponse `json:"tables"`
}

type TableStructureResponse struct {
	TableName string                    `json:"tableName"`
	Columns   []ColumnStructureResponse `json:"columns"`
}

type ColumnStructureResponse struct {
	ColumnName string `json:"columnName"`
	DataType   string `json:"dataType"`
}
