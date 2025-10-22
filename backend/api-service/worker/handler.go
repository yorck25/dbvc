package worker

import (
	"backend/connectors"
	"backend/core"
	"database/sql"
	"fmt"
	"strconv"
)

func HandleUpgrade(ctx *core.WebContext) error {
	projectID, err := strconv.Atoi(ctx.Request().Header.Get("Id"))
	if err != nil {
		return ctx.InternalError("Fail to upgrade: " + err.Error())
	}

	err = Upgrade(ctx, projectID)
	if err != nil {
		return ctx.InternalError("Fail to upgrade: " + err.Error())
	}

	return ctx.Sucsess()
}

func Upgrade(ctx *core.WebContext, id int) error {
	repo := NewRepository(ctx)

	project, err := repo.GetProjectDetails(id)
	if err != nil {
		return err
	}

	connectionTypeObj, err := repo.GetConnectionTypeObject(project.ConnectionType)
	if err != nil {
		return err
	}

	versionsSinceLatestRelease, err := repo.GetVersionsSinceLatestRelease(project.ID, 0)
	if err != nil {
		return err
	}

	connector, err := GetConnector(connectionTypeObj.Key)
	if err != nil {
		return err
	}

	db, err := connector.Connect()
	if err != nil {
		fmt.Println("Connection error:", err)
		return err
	}
	defer db.Close()

	_, err = GetDatabaseVersion(db, connector)
	if err != nil {
		return err
	}

	for _, version := range versionsSinceLatestRelease {
		fmt.Println("Execute Script: " + version.Up.Script)
		connector.ExecuteQuery(db, version.Up.Script)
	}

	return nil
}

func GetDatabaseVersion(db *sql.DB, connector connectors.DBConnector) (string, error) {
	var version string

	res, err := connector.ExecuteQuery(db, connector.GetVersionQuery())
	if err != nil {
		return version, err
	}

	if res.Next() {
		if err := res.Scan(&version); err != nil {
			return version, err
		}
	} else {
		return version, fmt.Errorf("no rows returned")
	}

	return version, nil
}
