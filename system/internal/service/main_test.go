// +build integration

package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/namsral/flag"
	"github.com/titpetric/factory"
	"go.uber.org/zap/zapcore"

	"github.com/crusttech/crust/internal/logger"
	systemMigrate "github.com/crusttech/crust/system/db"
)

func TestMain(m *testing.M) {
	logger.Init(zapcore.DebugLevel)

	dsn := ""
	flag.StringVar(&dsn, "db-dsn", "crust:crust@tcp(crust-db:3306)/crust?collation=utf8mb4_general_ci", "DSN for database connection")
	flag.Parse()

	factory.Database.Add("default", dsn)
	factory.Database.Add("system", dsn)

	db := factory.Database.MustGet()
	db.Profiler = &factory.Database.ProfilerStdout

	// migrate database schema
	if err := systemMigrate.Migrate(db); err != nil {
		fmt.Printf("Error running migrations: %+v\n", err)
		return
	}

	os.Exit(m.Run())
}
