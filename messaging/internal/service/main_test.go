// +build integration external

package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/namsral/flag"
	"github.com/titpetric/factory"
	"go.uber.org/zap/zapcore"

	"github.com/crusttech/crust/internal/logger"
	messagingMigrate "github.com/crusttech/crust/messaging/db"
	"github.com/crusttech/crust/messaging/internal/repository"
	systemMigrate "github.com/crusttech/crust/system/db"
	systemService "github.com/crusttech/crust/system/service"
)

type mockDB struct{}

func (mockDB) Transaction(callback func() error) error { return callback() }

func TestMain(m *testing.M) {
	logger.Init(zapcore.DebugLevel)

	dsn := ""
	new(repository.Flags).Init("messaging")
	flag.StringVar(&dsn, "db-dsn", "crust:crust@tcp(crust-db:3306)/crust?collation=utf8mb4_general_ci", "DSN for database connection")
	flag.Parse()

	factory.Database.Add("default", dsn)
	factory.Database.Add("messaging", dsn)
	factory.Database.Add("system", dsn)

	// @todo: add flags to configure the sql profiler in tests

	db := factory.Database.MustGet()
	db.Profiler = &factory.Database.ProfilerStdout

	dbSystem := factory.Database.MustGet("system")
	dbSystem.Profiler = &factory.Database.ProfilerStdout

	// migrate database schema
	if err := systemMigrate.Migrate(dbSystem); err != nil {
		fmt.Printf("Error running migrations: %+v\n", err)
		return
	}
	if err := messagingMigrate.Migrate(db); err != nil {
		fmt.Printf("Error running migrations: %+v\n", err)
		return
	}

	// clean up tables
	{
		for _, name := range []string{"sys_user", "sys_role", "sys_role_member", "sys_organisation"} {
			_, err := db.Exec("truncate " + name)
			if err != nil {
				panic("Error when clearing " + name + ": " + err.Error())
			}
		}
	}

	systemService.Init(context.Background())
	Init(context.Background())

	os.Exit(m.Run())
}
