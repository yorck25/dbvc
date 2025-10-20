package worker

import (
	"backend/connectors"
	"fmt"
)

func HandleUpgrade(connectionType string) error {
	var connector connectors.DBConnector

	switch connectionType {
	case "postgres":
		connector = connectors.PostgresConnector{
			ConnectionString: "host=localhost port=5432 user=postgres password=secret dbname=mydb sslmode=disable",
		}
	case "mysql":
		connector = connectors.MySQLConnector{
			ConnectionString: "user:password@tcp(localhost:3306)/mydb",
		}
	case "mssql":
		connector = connectors.MSSQLConnector{
			ConnectionString: "sqlserver://user:password@localhost:1433?database=mydb",
		}
	}

	db, err := connector.Connect()
	if err != nil {
		fmt.Println("Connection error:", err)
		return err
	}
	defer db.Close()

	return nil
}
