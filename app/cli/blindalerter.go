package cli

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(scheduledAt time.Duration, amount int)
}

type BlindAlerterFunc func(scheduledAt time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlertAt(scheduledAt time.Duration, amount int) {
	a(scheduledAt, amount)
}

func StdOutBlindAlerter(scheduledAt time.Duration, amount int) {
	time.AfterFunc(scheduledAt, func() {
		_, _ = fmt.Fprintf(os.Stdout, "Blind is now: %d\n", amount)
	})
}
