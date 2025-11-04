package common

import (
	"backend/connectors"
	"fmt"
)

func GetConnector(connectionType string, authData map[string]string) (connectors.DBConnector, error) {
	var connector connectors.DBConnector

	switch connectionType {
	case "psql":
		connectionString := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			authData["host"],
			authData["port"],
			authData["username"],
			authData["password"],
			authData["databaseName"],
		)

		connector = connectors.PostgresConnector{
			ConnectionString: connectionString,
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
