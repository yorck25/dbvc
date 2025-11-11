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
		connector = &connectors.PostgresConnector{
			Client:           nil,
			ConnectionString: connectionString,
		}

	case "mysql":
		connectionString := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			authData["username"],
			authData["password"],
			authData["host"],
			authData["port"],
			authData["databaseName"],
		)
		connector = &connectors.MySQLConnector{
			Client:           nil,
			ConnectionString: connectionString,
		}

	case "mssql":
		connectionString := fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s",
			authData["username"],
			authData["password"],
			authData["host"],
			authData["port"],
			authData["databaseName"],
		)
		connector = &connectors.MSSQLConnector{
			Client:           nil,
			ConnectionString: connectionString,
		}

	default:
		return nil, fmt.Errorf("unsupported connection type: %s", connectionType)
	}

	return connector, nil
}
