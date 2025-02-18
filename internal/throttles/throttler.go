package throttles

import (
	"time"

	interfaces "github.com/shoplineapp/captin/interfaces"
	log "github.com/sirupsen/logrus"
)

var tLogger = log.WithFields(log.Fields{"class": "Throttler"})

// Throttler - Event Throttler
type Throttler struct {
	interfaces.ThrottleInterface
	store interfaces.StoreInterface
}

// NewThrottler - Create new Throttler
func NewThrottler(store interfaces.StoreInterface) *Throttler {
	return &Throttler{
		store: store,
	}
}

// CanTrigger - Check if can trigger
func (t *Throttler) CanTrigger(id string, period time.Duration) (bool, time.Duration, error) {
	val, ok, duration, err := t.store.Get(id)

	if err != nil {
		return true, time.Duration(0), err
	}
	tLogger.WithFields(log.Fields{"value": val}).Debug("Check throttle value on CanTrigger")
	if !ok {
		tLogger.WithFields(log.Fields{"period": period}).Error("Unable to create throttle in store with period")
		t.store.Set(id, "1", period)
		return true, time.Duration(0), nil
	}

	return false, duration, nil
}
