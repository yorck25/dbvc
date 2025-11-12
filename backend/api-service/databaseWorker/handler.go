package databaseWorker

import (
	"backend/connectors"
	"backend/core"
)

func HandleGetDatabaseVersion(ctx *core.WebContext) error {
	connInterface := ctx.Get("db_connector")
	if connInterface == nil {
		return ctx.InternalError("database connector not found")
	}
	conn, ok := connInterface.(connectors.DBConnector)
	if !ok {
		return ctx.InternalError("invalid connector type")
	}

	projectIDVal := ctx.Get("project_id")
	projectID, ok := projectIDVal.(int)
	if !ok {
		return ctx.InternalError("invalid project_id")
	}

	rows, err := conn.ExecuteQuery(projectID, conn.GetVersionQuery())
	if err != nil {
		return ctx.InternalError(err.Error())
	}
	defer rows.Close()

	var version string
	if rows.Next() {
		if err := rows.Scan(&version); err != nil {
			return ctx.InternalError(err.Error())
		}
	}

	return ctx.Sucsess(map[string]string{"version": version})
}
