package battlenet

import (
	"time"
)

type Quota interface {
	// QPSAllotted returns the allotted QPS.
	QPSAllotted() int

	// QPSCurrent returns the current QPS.
	QPSCurrent() int

	// QuotaAlloted returns the alloted quota.
	QuotaAlloted() int

	// QuotaCurrent returns the current quota.
	QuotaCurrent() int

	// QuotaReset returns the time the quota will reset.
	QuotaReset() time.Time
}
