package worker

import (
	"backend/connectors"
	"fmt"
)

func GetConnector(connectionType string) (connectors.DBConnector, error) {
	var connector connectors.DBConnector

	switch connectionType {
	case "psql":
		connector = connectors.PostgresConnector{
			ConnectionString: "host=localhost port=5432 user=postgres password=secret dbname=test_db sslmode=disable",
		}
	case "mysql":
		connector = connectors.MySQLConnector{
			ConnectionString: "user:password@tcp(localhost:3306)/mydb",
		}
	case "mssql":
		connector = connectors.MSSQLConnector{
			ConnectionString: "sqlserver://user:password@localhost:1433?database=mydb",
		}
	default:
		return connector, fmt.Errorf("unsupported connection type")
	}

	return connector, nil
}

func CreateUpgradeScript() {
}

func CreateDowngradeScript() {

}
