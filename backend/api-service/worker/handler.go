package worker

import (
	"backend/common"
	"backend/connectors"
	"backend/core"
	"backend/projects"
	"backend/release"
	"database/sql"
	"fmt"
	"strconv"
)

func HandleGetConnectionTypes(ctx *core.WebContext) error {
	_, err := ctx.GetUserId()
	if err != nil {
		return err
	}

	repo := NewRepository(ctx)

	connectionTypes, err := repo.GetConnectionTypes()
	if err != nil {
		return ctx.InternalError("Fail to get connection types: " + err.Error())
	}

	return ctx.Sucsess(connectionTypes)
}

func HandleGetDatabaseStructure(ctx *core.WebContext) error {
	_, err := ctx.GetUserId()
	if err != nil {
		return err
	}

	var dba projects.DatabaseAuth

	if err := ctx.Bind(&dba); err != nil {
		return ctx.BadRequest("invalid input")
	}

	dbTypeVal, ok := dba.DatabaseAuth["type"]
	if !ok || dbTypeVal == "" {
		return ctx.BadRequest("database type is required")
	}

	_, connector, err := CreateDatabaseClient(dbTypeVal, dba.DatabaseAuth)
	if err != nil {
		fmt.Println("test")
		return ctx.InternalError(err.Error())
	}

	structure, err := connector.GetDatabaseStructure()
	if err != nil {
		fmt.Println("test 22")
		return ctx.InternalError(err.Error())
	}

	return ctx.Sucsess(structure)
}

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

	releaseRepo := release.NewRepository(ctx)
	release, err := releaseRepo.GetLatestReleasesForProject(project.ID)
	if err != nil {
		return err
	}

	versionsSinceLatestRelease, err := repo.GetVersionsSinceLatestRelease(project.ID, release.CurrentVersion)
	if err != nil {
		return err
	}

	testParams := map[string]string{
		"host": "127.0.0.1", "port": "5432", "user": "postgres_admin", "password": "test1234", "database": "test",
	}

	connector, err := common.GetConnector(connectionTypeObj.Key, testParams)
	if err != nil {
		return err
	}

	db, err := connector.Connect()
	if err != nil {
		fmt.Println("Connection error:", err)
		return err
	}

	_, err = GetDatabaseVersion(db, connector)
	if err != nil {
		return err
	}

	for _, version := range versionsSinceLatestRelease {
		fmt.Println("Execute Script: " + version.Up.Script)
		connector.ExecuteQuery(version.Up.Script)
	}

	return nil
}

func GetDatabaseVersion(db *sql.DB, connector connectors.DBConnector) (string, error) {
	var version string

	res, err := connector.ExecuteQuery(connector.GetVersionQuery())
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
