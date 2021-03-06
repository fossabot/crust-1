package db

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/titpetric/factory"
	"go.uber.org/zap"

	"github.com/crusttech/crust/internal/logger"
)

const (
	maxTries = 100
	delay    = 5 * time.Second
	timeout  = 1 * time.Minute
)

func TryToConnect(ctx context.Context, name, dsn, profiler string) (db *factory.DB, err error) {
	factory.Database.Add(name, dsn)

	var (
		connErrCh = make(chan error, 1)
		log       = logger.Default().With(zap.String("name", name), zap.String("dsn", dsn))
	)

	defer close(connErrCh)

	log.Debug("connecting to the database")

	go func() {
		var (
			try = 0
		)

		for {
			try++

			if maxTries <= try {
				err = errors.Errorf("could not connect to %q, in %d tries", name, try)
				return
			}

			db, err = factory.Database.Get(name)
			if err != nil {
				log.Warn(
					"could not connect",
					zap.Error(err),
					zap.Int("try", try),
					zap.Float64("delay", delay.Seconds()),
				)

				select {
				case <-ctx.Done():
					// Forced break
					break
				case <-time.After(delay):
					// Wait before next try
					continue
				}
			}

			// Connected
			break

		}

		connErrCh <- err
	}()

	select {
	case err = <-connErrCh:
		break
	case <-time.After(timeout):
		// Wait before next try
		return nil, errors.Errorf("db init for %q timedout", name)
	case <-ctx.Done():
		return nil, errors.Errorf("db connection for %q cancelled", name)
	}

	// @todo: profiling as an external service?
	switch profiler {
	case "stdout":
		db.Profiler = &factory.Database.ProfilerStdout
	case "logger":
		db.Profiler = ZapProfiler(logger.Default().
			Named("database").
			// Skip 3 levels in call stack to get to the actual function used
			WithOptions(zap.AddCallerSkip(3)).
			With(zap.String("connection", name)),
		)
	default:
		log.Info("no database query profiler selected")
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
