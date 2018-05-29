package cmdTable

import (
	"github.com/procroner/ssh-run/core/coreError"
	"fmt"
	"strings"
	"github.com/procroner/ssh-run/core/coreJob"
	"github.com/procroner/ssh-run/core/coreServer"
	"github.com/procroner/ssh-run/core/coreLog"
	"errors"
)

func MigrateTables(tables string) {
	if tables == "" {
		coreServer.Migrate()
		coreJob.Migrate()
		coreLog.Migrate()
	} else {
		tablesSlice := strings.Split(tables, ",")
		for _, table := range tablesSlice {
			switch table {
			case "servers", "server":
				coreServer.Migrate()
			case "jobs", "job":
				coreJob.Migrate()
			case "logs", "log":
				coreLog.Migrate()
			default:
				coreError.HandleError("db", errors.New(fmt.Sprintf("DB: table %s is not needed.", table)))
			}
		}
	}
}
