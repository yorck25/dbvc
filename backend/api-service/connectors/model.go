package connectors

import (
	"database/sql"
)

type DBConnector interface {
	Connect() (*sql.DB, error)
	Disconnect() error

	ExecuteQuery(string) (*sql.Rows, error)
	GetVersionQuery() string
	GetDatabaseStructure() (*DatabaseStructureResponse, error)
}

type PostgresConnector struct {
	Client           *sql.DB
	ConnectionString string
}

type MySQLConnector struct {
	Client           *sql.DB
	ConnectionString string
}

type MSSQLConnector struct {
	Client           *sql.DB
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
