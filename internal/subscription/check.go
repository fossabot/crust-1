package subscription

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/crusttech/permit/pkg/permit"

	"github.com/crusttech/crust/internal/logger"
)

const (
	// How many times we try to get the licence
	checkMaxTries = 10

	// Number of seconds between each try (x number of tries)
	checkTryDelay = 10

	// Timeout in seconds
	checkTimeout = 30
)

// Check for subscription
func Check(ctx context.Context) (err error) {
	var (
		done context.CancelFunc
		p    *permit.Permit
		try  = 1
		in   = permit.Permit{
			Key:    flags.subscription.Key,
			Domain: flags.subscription.Domain,
		}
	)

	// Do not collect stats on local domains.
	// if p.Domain != "local.crust.tech" {
	// @todo collect & pass attributes (no of users....) to be validated by permit.crust.tech subscription server.
	in.Attributes = map[string]int{
		"messaging.enabled": 1,
		// "messaging.max-public-channels": 1,
		// "messaging.max-messages": 1,
		// "messaging.max-users": 1,
		// "messaging.max-private-channels": 1,

		"system.enabled": 1,
		// "system.max-organisations": 1,
		// "system.max-users": 1,
		// "system.max-teams": 1,

		"compose.enabled": 1,
		// "compose.max-modules": 1,
		// "compose.max-pages": 1,
		// "compose.max-triggers": 1,
		// "compose.max-users": 1,
		// "compose.max-namespaces": 1,
		// "compose.max-charts": 1,
	}
	// }

	for {
		ctx, done = context.WithTimeout(ctx, time.Second*checkTimeout)
		defer done()

		p, err = permit.Check(ctx, in)
		if err == nil {
			break
		}

		if err != nil {
			logger.Default().Warn(
				"unable to check for subscription",
				zap.Error(err),
				zap.Int("try", try),
				zap.Int("max", checkMaxTries),
			)
		}

		if try >= checkMaxTries {
			return errors.Wrap(err, "unable to check for subscription")
		}

		time.Sleep(time.Second * checkTryDelay * checkTryDelay)

		try++
	}

	if !p.IsValid() {
		return err
	} else if p.Domain != flags.subscription.Domain {
		return errors.Errorf("subscription domains do not match (%s <> %s)", p.Domain, flags.subscription.Domain)
	}

	return nil
}
