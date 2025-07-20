package poker

import (
	"fmt"
	"io"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(scheduledAt time.Duration, amount int, to io.Writer)
}

type BlindAlerterFunc func(scheduledAt time.Duration, amount int, to io.Writer)

func (a BlindAlerterFunc) ScheduleAlertAt(scheduledAt time.Duration, amount int, to io.Writer) {
	a(scheduledAt, amount, to)
}

func Alerter(scheduledAt time.Duration, amount int, to io.Writer) {
	time.AfterFunc(scheduledAt, func() {
		_, _ = fmt.Fprintf(to, "Blind is now: %d\n", amount)
	})
}
