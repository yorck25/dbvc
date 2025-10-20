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
	fmt.Println("Executing query on PostgreSQL...")
	return db.Query(query)
}
