package connectors

import "database/sql"

type DBConnector interface {
	Connect() (*sql.DB, error)
	ExecuteQuery(db *sql.DB, query string) (*sql.Rows, error)
}

type PostgresConnector struct {
	ConnectionString string
}

type MySQLConnector struct {
	ConnectionString string
}

type MSSQLConnector struct {
	ConnectionString string
}
