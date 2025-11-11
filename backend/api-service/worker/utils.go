package worker

import (
	"backend/common"
	"backend/connectors"
	"database/sql"
	"fmt"
)

func CreateUpgradeScript() {
}

func CreateDowngradeScript() {

}

func CreateDatabaseClient(dbTypeVal string, databaseAuth map[string]string) (*sql.DB, connectors.DBConnector, error) {

	connector, err := common.GetConnector(dbTypeVal, databaseAuth)
	if err != nil {
		return nil, nil, err
	}

	db, err := connector.Connect()
	if err != nil {
		fmt.Println("Connection error:", err)
		return nil, nil, err
	}

	return db, connector, nil
}
